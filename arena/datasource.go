package arena

import (
	"encoding/json"
	"os"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena/thcurvegenresponsev1"
	"github.com/klauspost/compress/arena/thfooitem"
	"github.com/klauspost/compress/arena/thvitalstemplate"
	"github.com/tinylib/msgp/msgp"

	"google.golang.org/protobuf/proto"
)

type DatasourceEntry struct {
	Item         msgp.Encodable
	ThriftItem   thrift.TStruct
	ProtobufItem proto.Message

	HambaAvroSchema avro.Schema

	Bytes                []byte                //for deserialization tests
	NewEmptyItem         func() msgp.Decodable //for deserialization tests
	NewEmptyThriftItem   func() thrift.TStruct //for deserialization tests
	NewEmptyProtobufItem func() proto.Message  //for deserialization tests
}

var MainDatasource = []DatasourceEntry{}

func InitializeMainDatasource() {
	sampleVitalsTemplatePatched := &VitalsTemplate{}
	sampleCurveGenReplyV1_fromCurveGenSrv_toManikin := &CurveGenReplyV1{}
	sampleCurveGenReplyV1_fromCurveGenSrv_toPatientMonitor := &CurveGenReplyV1{}

	err := json.Unmarshal(
		[]byte(Curvegenreplyv1___from_curvegensrv_for_manikin___json_string),
		sampleCurveGenReplyV1_fromCurveGenSrv_toManikin,
	)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(
		[]byte(Curvegenreplyv1___from_curvegensrv_for_patientmonitor___json_string),
		sampleCurveGenReplyV1_fromCurveGenSrv_toPatientMonitor,
	)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(
		[]byte(Vitalset_patched___json_string),
		sampleVitalsTemplatePatched,
	)
	if err != nil {
		panic(err)
	}

	goHambaAvroFooItemSchema, err := os.ReadFile("../avfooitem.fixedmanually.avsc") // intentionally avoided 'avfooitem.avsc' because its problematic
	if err != nil {
		panic(err)
	}

	goHambaAvroCurveGenResponseV1Schema, err := os.ReadFile("../avcurvegenresponsev1.fixedmanually.avsc") // intentionally avoided 'avcurvegenresponsev1.avsc' because its problematic
	if err != nil {
		panic(err)
	}

	goHambaAvroVitalSetPatchedSchema, err := os.ReadFile("../avvitalsetpatched.fixedmanually.avsc") // intentionally avoided 'avvitalsetpatched.avsc' because its problematic
	if err != nil {
		panic(err)
	}

	MainDatasource = append(
		MainDatasource,
		DatasourceEntry{
			Item:         sampleCurveGenReplyV1_fromCurveGenSrv_toManikin,
			ThriftItem:   ConvertCurvegenResponseV1ToThrift(sampleCurveGenReplyV1_fromCurveGenSrv_toPatientMonitor),
			ProtobufItem: ConvertCurvegenResponseV1ToProtobuf(sampleCurveGenReplyV1_fromCurveGenSrv_toPatientMonitor),

			NewEmptyItem:         func() msgp.Decodable { return &CurveGenReplyV1{} },
			NewEmptyThriftItem:   func() thrift.TStruct { return thcurvegenresponsev1.NewTHCurveGenReplyV1() },
			NewEmptyProtobufItem: func() proto.Message { return &PBCurveGenReplyV1{} },

			HambaAvroSchema: avro.MustParse(string(goHambaAvroCurveGenResponseV1Schema)),
		},
	)

	MainDatasource = append(
		MainDatasource,
		DatasourceEntry{
			Item:         sampleCurveGenReplyV1_fromCurveGenSrv_toPatientMonitor,
			ThriftItem:   ConvertCurvegenResponseV1ToThrift(sampleCurveGenReplyV1_fromCurveGenSrv_toPatientMonitor),
			ProtobufItem: ConvertCurvegenResponseV1ToProtobuf(sampleCurveGenReplyV1_fromCurveGenSrv_toPatientMonitor),

			NewEmptyItem:         func() msgp.Decodable { return &CurveGenReplyV1{} },
			NewEmptyThriftItem:   func() thrift.TStruct { return thcurvegenresponsev1.NewTHCurveGenReplyV1() },
			NewEmptyProtobufItem: func() proto.Message { return &PBCurveGenReplyV1{} },

			HambaAvroSchema: avro.MustParse(string(goHambaAvroCurveGenResponseV1Schema)),
		},
	)

	MainDatasource = append(
		MainDatasource,
		DatasourceEntry{
			Item:         sampleVitalsTemplatePatched,
			ThriftItem:   ConvertVitalsTemplateToThrift(sampleVitalsTemplatePatched),
			ProtobufItem: ConvertVitalsTemplateToProtobuf(sampleVitalsTemplatePatched),

			NewEmptyItem:         func() msgp.Decodable { return &VitalsTemplate{} },
			NewEmptyThriftItem:   func() thrift.TStruct { return thvitalstemplate.NewTHVitalsTemplate() },
			NewEmptyProtobufItem: func() proto.Message { return &PBVitalsTemplate{} },

			HambaAvroSchema: avro.MustParse(string(goHambaAvroVitalSetPatchedSchema)),
		},
	)

	fooitem := &FooItem{
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
	}

	MainDatasource = append(
		MainDatasource,
		DatasourceEntry{
			Item:         fooitem,
			ThriftItem:   ConvertFooItemToThrift(fooitem),
			ProtobufItem: ConvertFooItemToProtobuf(fooitem),

			NewEmptyItem:         func() msgp.Decodable { return &FooItem{} },
			NewEmptyThriftItem:   func() thrift.TStruct { return thfooitem.NewTHFooItem() },
			NewEmptyProtobufItem: func() proto.Message { return &PBFooItem{} },

			HambaAvroSchema: avro.MustParse(string(goHambaAvroFooItemSchema)),
		},
	)
}

const TimeFormat = "2006-01-02T15:04:05.000Z"

func parseTimeString(input string) time.Time {

	t, err := time.Parse(TimeFormat, input)
	if err != nil {
		return time.Time{}
	}

	return t
}
