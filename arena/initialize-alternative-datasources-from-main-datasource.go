package arena

import (
	"context"
	"encoding/json"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/fxamacker/cbor/v2"
	"github.com/vmihailenco/msgpack/v5"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"

	"github.com/klauspost/compress/arena/thfooitem"
)

type serializedDataSources struct {
	Json         [][]byte
	Cbor         [][]byte
	MessagePack  [][]byte
	Bson         [][]byte
	Protobuf     [][]byte
	ThriftBinary [][]byte
}

type datasourcesForIDLMechanisms struct {
	Protobuf []*PBFooItem
	Thrift   []*thfooitem.THFooItem
}

var SerializedDataSources = serializedDataSources{}
var SpecialDatasourcesForIDLMechanisms = datasourcesForIDLMechanisms{}

func InitializeAlternativeDatasourcesFromMainDatasource() {
	thriftSerializer := thrift.NewTSerializer()

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

		//bson
		bsonBytes, err := bson.Marshal(x)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.Bson = append(SerializedDataSources.Bson, bsonBytes)

		// thrift binary
		thFooItem := ConvertFooItemToTHFooItem(x)
		SpecialDatasourcesForIDLMechanisms.Thrift = append(SpecialDatasourcesForIDLMechanisms.Thrift, &thFooItem)

		thriftBytes, err := thriftSerializer.Write(context.TODO(), &thFooItem)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.ThriftBinary = append(SerializedDataSources.ThriftBinary, thriftBytes)

		//protobuf
		pbFooItem := ConvertFooItemToPBFooItem(x)
		SpecialDatasourcesForIDLMechanisms.Protobuf = append(SpecialDatasourcesForIDLMechanisms.Protobuf, &pbFooItem)

		protobufBytes, err := proto.Marshal(&pbFooItem)
		if err != nil {
			panic(err)
		}
		SerializedDataSources.Protobuf = append(SerializedDataSources.Protobuf, protobufBytes)
	}
}
