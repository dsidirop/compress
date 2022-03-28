package arena

import (
	"fmt"
	"time"

	"github.com/klauspost/compress/arena/thsimeventregistereventcmd"
)

func ConvertSimEventRegisterEventCmdToThrift(x *SimEventRegisterEventCmd) *thsimeventregistereventcmd.THSimEventRegisterEventCmd {
	if x == nil {
		return nil
	}

	timestampBytes, _ := x.TimeStamp.UTC().MarshalText()

	results := thsimeventregistereventcmd.NewTHSimEventRegisterEventCmd()
	results.Tenant = x.Tenant
	results.EventID = x.EventID
	results.Arguments = convertSimEventRegisterEventCmdToThrift_arguments(x.EventID, x.Arguments)
	results.TimeStamp = thsimeventregistereventcmd.Timestamp(timestampBytes)
	results.ApiVersion = x.ApiVersion
	results.SimulatorName = x.SimulatorName

	return results
}

func convertSimEventRegisterEventCmdToThrift_arguments(eventID string, arguments map[string]interface{}) *thsimeventregistereventcmd.THSimEventArguments {
	if arguments == nil {
		return nil
	}

	thSimEventPacingArgumentsWrapper := thsimeventregistereventcmd.NewTHSimEventArguments()

	if eventID == "psri.event.PacingOccurred" {
		timestamp := time.UnixMilli(int64(arguments["time_stamp"].(float64)))
		timestampBytes, _ := timestamp.UTC().MarshalText()

		thSimEventPacingArgumentsWrapper.PacingArguments = thsimeventregistereventcmd.NewTHPacingArguments()
		thSimEventPacingArgumentsWrapper.PacingArguments.Capture = arguments["capture"].(bool)
		thSimEventPacingArgumentsWrapper.PacingArguments.TimeStamp = thsimeventregistereventcmd.Timestamp(timestampBytes)
		thSimEventPacingArgumentsWrapper.PacingArguments.MeasuredCurrent = int32(arguments["measured_current"].(float64))

		return thSimEventPacingArgumentsWrapper
	}

	panic(fmt.Sprintf("unknown type '%T'", arguments))
}
