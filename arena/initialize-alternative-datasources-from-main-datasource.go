package arena

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/fxamacker/cbor/v2"
	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena/thfooitem"
	"github.com/klauspost/compress/s2"
	"github.com/klauspost/compress/zstd"
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

type serializedAndCompressedDataSources struct {
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

type compressionTestCase struct {
	Desc                  string
	CompressionCallback   func([]byte) ([]byte, error)
	DecompressionCallback func([]byte) ([]byte, error)
}

var Schemas = schemas{}
var SerializedDataSources = serializedDataSources{}
var AllCompressionTestCases = []compressionTestCase{}
var SpecialDatasourcesForIDLMechanisms = datasourcesForIDLMechanisms{}
var SerializedAndCompressedDataSources = serializedAndCompressedDataSources{}

func InitTestProvisions() {
	InitCompressionTestCases()

	InitIDLSchemas()                                     //   order
	InitializeAlternativeDatasourcesFromMainDatasource() //   order
}

func InitCompressionTestCases() {
	AllCompressionTestCases = []compressionTestCase{
		func() compressionTestCase {
			return compressionTestCase{
				Desc: "ZLib",
				CompressionCallback: func(rawBytes []byte) ([]byte, error) {
					byteBuffer := bytes.Buffer{}
					w := zlib.NewWriter(&byteBuffer)
					_, err := w.Write(rawBytes)
					w.Close()

					return byteBuffer.Bytes(), err
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					byteBuffer := bytes.NewReader(compressedBytes)
					zlibReader, err := zlib.NewReader(byteBuffer)
					if err != nil {
						return nil, err
					}

					buf := bytes.Buffer{}
					buf.ReadFrom(zlibReader)
					zlibReader.Close()

					return buf.Bytes(), err
				},
			}
		}(),
		func() compressionTestCase {
			encoder, err := zstd.NewWriter(nil)
			if err != nil {
				panic(err)
			}

			decoder, err := zstd.NewReader(nil)
			if err != nil {
				panic(err)
			}

			return compressionTestCase{
				Desc: "ZStandard",
				CompressionCallback: func(rawBytes []byte) ([]byte, error) {
					compressedBytes := encoder.EncodeAll(rawBytes, nil)
					//encoder.Close() //do we need this?

					return compressedBytes, nil
				},
				DecompressionCallback: func(rawBytes []byte) ([]byte, error) {
					decompressedBytes, err := decoder.DecodeAll(rawBytes, nil)
					//decoder.Close() //do we need this?

					return decompressedBytes, err
				},
			}
		}(),
		func() compressionTestCase {
			return compressionTestCase{
				Desc: "S2",
				CompressionCallback: func(rawBytes []byte) ([]byte, error) {
					destination := &bytes.Buffer{}

					encoder := s2.NewWriter(destination)
					_, err := io.Copy(encoder, bytes.NewReader(rawBytes))
					encoder.Close()

					return destination.Bytes(), err
				},
				DecompressionCallback: func(rawBytes []byte) ([]byte, error) {
					reader := bytes.NewReader(rawBytes)

					decoder := s2.NewReader(reader)
					decompressedResultBuffer := &bytes.Buffer{}
					_, err := io.Copy(decompressedResultBuffer, decoder)

					return decompressedResultBuffer.Bytes(), err
				},
			}
		}(),
	}
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
