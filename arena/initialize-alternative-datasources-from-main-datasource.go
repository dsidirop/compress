package arena

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/fxamacker/cbor/v2"
	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena/thfooitem"
	"github.com/tinylib/msgp/msgp"
	"github.com/vmihailenco/msgpack/v5"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
)

type serializedDataSources struct {
	Json          [][]byte
	Cbor          [][]byte
	MessagePack   [][]byte
	Msgp          [][]byte
	Bson          [][]byte
	Protobuf      [][]byte
	ThriftBinary  [][]byte
	ThriftCompact [][]byte
	GoHambaAvro   [][]byte
}

type datasourcesForIDLMechanisms struct {
	Protobuf []*PBFooItem
	Thrift   []*thfooitem.THFooItem
}

type schemas struct {
	GoHambaAvro avro.Schema
}

var Schemas = schemas{}
var SerializedDataSources = serializedDataSources{}
var SpecialDatasourcesForIDLMechanisms = datasourcesForIDLMechanisms{}

func InitTestProvisions() {
	InitIDLSchemas()                                     //   order
	InitializeAlternativeDatasourcesFromMainDatasource() //   order
}

func InitIDLSchemas() {
	goAvroSchema, err := os.ReadFile("../avfooitem.fixedmanually.avsc") // intentionally avoid 'avfooitem.avsc' because
	if err != nil {
		log.Fatal(err)
	}
	Schemas.GoHambaAvro = avro.MustParse(string(goAvroSchema))
}

func InitializeAlternativeDatasourcesFromMainDatasource() {
	thriftBinarySerializer := thrift.NewTSerializer()
	thriftCompactSerializer := NewThriftCompactSerializer()

	datasourceArrayLength := len(Datasource)
	for i := 0; i < datasourceArrayLength; i++ {
		x := Datasource[i]

		//json
		jsonBytes, err := json.Marshal(x)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.Json = append(SerializedDataSources.Json, jsonBytes)

		//cbor
		cborBytes, err := cbor.Marshal(x)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.Cbor = append(SerializedDataSources.Cbor, cborBytes)

		//messagepack
		messagePackBytes, err := msgpack.Marshal(x)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.MessagePack = append(SerializedDataSources.MessagePack, messagePackBytes)

		//msgp
		buf := bytes.Buffer{}
		err = msgp.Encode(&buf, &x)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.Msgp = append(SerializedDataSources.Msgp, buf.Bytes())

		//bson
		bsonBytes, err := bson.Marshal(x)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.Bson = append(SerializedDataSources.Bson, bsonBytes)

		//thrift
		thFooItem := ConvertFooItemToTHFooItem(x)
		SpecialDatasourcesForIDLMechanisms.Thrift = append(SpecialDatasourcesForIDLMechanisms.Thrift, &thFooItem)

		//thrift-binary
		thriftBinaryBytes, err := thriftBinarySerializer.Write(context.TODO(), &thFooItem)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.ThriftBinary = append(SerializedDataSources.ThriftBinary, thriftBinaryBytes)

		//thrift-compact
		thriftCompactBytes, err := thriftCompactSerializer.Write(context.TODO(), &thFooItem)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.ThriftCompact = append(SerializedDataSources.ThriftCompact, thriftCompactBytes)

		//protobuf
		pbFooItem := ConvertFooItemToPBFooItem(x)
		SpecialDatasourcesForIDLMechanisms.Protobuf = append(SpecialDatasourcesForIDLMechanisms.Protobuf, &pbFooItem)

		protobufBytes, err := proto.Marshal(&pbFooItem)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.Protobuf = append(SerializedDataSources.Protobuf, protobufBytes)

		//goavro
		gohambaAvroBytes, err := avro.Marshal(Schemas.GoHambaAvro, &x)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.GoHambaAvro = append(SerializedDataSources.GoHambaAvro, gohambaAvroBytes)
	}
}
