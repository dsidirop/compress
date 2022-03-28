package arena

import (
	"fmt"
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertSimEventRegisterEventCmdToProtobuf(x *SimEventRegisterEventCmd) *PBSimEventRegisterEventCmd {
	if x == nil {
		return nil
	}

	results := &PBSimEventRegisterEventCmd{}
	results.Tenant = x.Tenant
	results.EventID = x.EventID
	results.Arguments = convertSimEventRegisterEventCmdToProtobuf_arguments(x.EventID, x.Arguments)
	results.TimeStamp = timestamppb.New(x.TimeStamp)
	results.ApiVersion = x.ApiVersion
	results.SimulatorName = x.SimulatorName

	return results
}

func convertSimEventRegisterEventCmdToProtobuf_arguments(eventID string, arguments map[string]interface{}) isPBSimEventRegisterEventCmd_Arguments {
	if arguments == nil {
		return nil
	}

	pbSimEventPacingArgumentsWrapper := &PBSimEventRegisterEventCmd_PacingArguments{}

	if eventID == "psri.event.PacingOccurred" {
		timestamp := time.UnixMilli(int64(arguments["time_stamp"].(float64)))

		pbSimEventPacingArgumentsWrapper.PacingArguments = &PBPacingArguments{}
		pbSimEventPacingArgumentsWrapper.PacingArguments.Capture = arguments["capture"].(bool)
		pbSimEventPacingArgumentsWrapper.PacingArguments.TimeStamp = timestamppb.New(timestamp)
		pbSimEventPacingArgumentsWrapper.PacingArguments.MeasuredCurrent = int32(arguments["measured_current"].(float64))

		return pbSimEventPacingArgumentsWrapper
	}

	panic(fmt.Sprintf("unknown type '%T'", arguments))
}
