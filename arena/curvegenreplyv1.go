package arena

type CurveGenReplyV1 struct {
	ClientId string               `    json:"client_id,omitempty"     msg:"client_id"            ` // For response message routing purposes
	Spec     CurveSpecificationV1 `    json:"spec"                    msg:"spec"                 `
	Lead1    []int64              `    json:"ECG_LEAD_1,omitempty"    msg:"ECG_LEAD_1"           `
	Lead2    []int64              `    json:"ECG_LEAD_2,omitempty"    msg:"ECG_LEAD_2"           `
	Lead3    []int64              `    json:"ECG_LEAD_3,omitempty"    msg:"ECG_LEAD_3"           `
	Lead4    []int64              `    json:"ECG_LEAD_4,omitempty"    msg:"ECG_LEAD_4"           `
	Lead5    []int64              `    json:"ECG_LEAD_5,omitempty"    msg:"ECG_LEAD_5"           `
	Lead6    []int64              `    json:"ECG_LEAD_6,omitempty"    msg:"ECG_LEAD_6"           `
	Lead7    []int64              `    json:"ECG_LEAD_7,omitempty"    msg:"ECG_LEAD_7"           `
	Lead8    []int64              `    json:"ECG_LEAD_8,omitempty"    msg:"ECG_LEAD_8"           `
	Lead9    []int64              `    json:"ECG_LEAD_9,omitempty"    msg:"ECG_LEAD_9"           `
	Lead10   []int64              `    json:"ECG_LEAD_10,omitempty"   msg:"ECG_LEAD_10"          `
	Lead11   []int64              `    json:"ECG_LEAD_11,omitempty"   msg:"ECG_LEAD_11"          `
	Lead12   []int64              `    json:"ECG_LEAD_12,omitempty"   msg:"ECG_LEAD_12"          `
	Abp      []int64              `    json:"ABP,omitempty"           msg:"ABP"                  `
	Cvp      []int64              `    json:"CVP,omitempty"           msg:"CVP"                  `
	Pap      []int64              `    json:"PAP,omitempty"           msg:"PAP"                  `
	Spo2     []int64              `    json:"SPO2,omitempty"          msg:"SPO2"                 `
	Wp       []int64              `    json:"WP,omitempty"            msg:"WP"                   ` // Time limited usage, shall only be shown when wedge inflated in pulmonary artery (very time limited)
	Tags     []Tag                `    json:"TAGS,omitempty"          msg:"TAGS"                 `
}

type CurveSpecificationV1 struct {
	SimulatorName  string `            json:"simulator_name"          msg:"simulator_name"       `
	Tenant         string `            json:"tenant"                  msg:"tenant"               `
	StartTime      int64  `            json:"start_time"              msg:"start_time"           ` // Time of first sample to be generated
	EndTime        int64  `            json:"end_time"                msg:"end_time"             ` // No sample after this time shall be included, end_time sample may not be included
	SampleInterval int64  `            json:"sample_interval"         msg:"sample_interval"      `
	CurveTypes     int64  `            json:"curve_types"             msg:"curve_types"          ` // Bitmapped according to CurveTypes enum
}

type Tag struct {
	TagTime int64 `                    json:"tm"                      msg:"tm"                   `
	TagType int64 `                    json:"tp"                      msg:"tp"                   `
}
