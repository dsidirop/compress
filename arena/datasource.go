package arena

import (
	"encoding/json"
	"time"

	"github.com/tinylib/msgp/msgp"
)

const TimeFormat = "2006-01-02T15:04:05.000Z"

func parseTimeString(input string) time.Time {

	t, err := time.Parse(TimeFormat, input)
	if err != nil {
		return time.Time{}
	}

	return t
}

var Datasource = []msgp.Encodable{}

func InitializeMainDatasource() {
	sampleCurveGenReplyV1_fromCurveGenSrv_toManikin := CurveGenReplyV1{}

	err := json.Unmarshal([]byte(Curvegenreplyv1___from_curvegensrv_for_manikin___json_string), &sampleCurveGenReplyV1_fromCurveGenSrv_toManikin)
	if err != nil {
		panic(err)
	}

	sampleCurveGenReplyV1_fromCurveGenSrv_toPatientMonitor := CurveGenReplyV1{}

	err = json.Unmarshal([]byte(Curvegenreplyv1___from_curvegensrv_for_patientmonitor___json_string), &sampleCurveGenReplyV1_fromCurveGenSrv_toPatientMonitor)
	if err != nil {
		panic(err)
	}

	Datasource = append(
		Datasource,
		&sampleCurveGenReplyV1_fromCurveGenSrv_toManikin,
	)

	Datasource = append(
		Datasource,
		&sampleCurveGenReplyV1_fromCurveGenSrv_toPatientMonitor,
	)

	Datasource = append(
		Datasource,
		&FooItem{
			ID:    "620e7d1f89c0231fc95854d8",
			Email: "latasha_hanson@digirang.mf",
			Roles: []string{
				"owner",
				"member",
			},
			APIKey: "790c6327-f9a6-487b-8cad-6a2ede14c4e5",
			Profile: FooProfile{
				Dob:     "1988-01-12",
				Name:    "Latasha Hanson",
				About:   "Ad aliqua ullamco nulla officia laborum do nulla et laboris nisi duis nisi consectetur. Do occaecat labore quis nulla pariatur non.",
				Address: "70 McKibben Street, Omar, Connecticut",
				Company: "Digirang",
				Location: FooLocation{
					Lat:  22.344456,
					Long: -79.955604,
				},
			},
			Username:  "latasha88",
			CreatedAt: parseTimeString("2013-11-22T01:00:35.839Z"),
			UpdatedAt: parseTimeString("2013-11-23T01:00:35.839Z"),
		},
	)
}
