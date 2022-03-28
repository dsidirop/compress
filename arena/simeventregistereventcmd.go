package arena

import "time"

type SimEventRegisterEventCmd struct {
	ApiVersion    string                 `     json:"api_version,omitempty"       msg:"api_version,omitempty"       bson:"api_version,omitempty"      `
	SimulatorName string                 `     json:"simulator_name,omitempty"    msg:"simulator_name,omitempty"    bson:"simulator_name,omitempty"   `
	Tenant        string                 `     json:"tenant,omitempty"            msg:"tenant,omitempty"            bson:"tenant,omitempty"           `
	EventID       string                 `     json:"event_id,omitempty"          msg:"event_id,omitempty"          bson:"event_id,omitempty"         `
	TimeStamp     time.Time              `     json:"time_stamp,omitempty"        msg:"time_stamp,omitempty"        bson:"time_stamp,omitempty"       `
	Arguments     map[string]interface{} `     json:"arguments,omitempty"         msg:"arguments,omitempty"         bson:"arguments,omitempty"        ` // SimEventPacingArguments and other types here
}

type SimEventPacingArguments struct {
	TimeStamp       time.Time `     json:"time_stamp,omitempty"        msg:"time_stamp,omitempty"        bson:"time_stamp,omitempty"       `
	MeasuredCurrent int32     `     json:"measured_current,omitempty"  msg:"measured_current,omitempty"  bson:"measured_current,omitempty" `
	Capture         bool      `     json:"capture,omitempty"           msg:"capture,omitempty"           bson:"capture,omitempty"          `
}
