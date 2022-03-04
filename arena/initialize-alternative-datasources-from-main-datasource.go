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
	"github.com/klauspost/compress/flate"
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
				CompressionCallback: func(uncompressedRawBytes []byte) ([]byte, error) {
					compressedBytesBufferWriter := &bytes.Buffer{}

					zlibCompressor := zlib.NewWriter(compressedBytesBufferWriter)
					_, err := zlibCompressor.Write(uncompressedRawBytes)
					if err != nil {
						return nil, err
					}

					zlibCompressor.Close() //dont use defer   it wont work

					return compressedBytesBufferWriter.Bytes(), err
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					compressedInputBuffer := bytes.NewReader(compressedBytes)
					compressedOutputBuffer := bytes.Buffer{}

					zlibCompressedInputReader, err := zlib.NewReader(compressedInputBuffer)
					if err != nil {
						return nil, err
					}

					_, err = compressedOutputBuffer.ReadFrom(zlibCompressedInputReader)
					zlibCompressedInputReader.Close()

					if err != nil {
						return nil, err
					}

					return compressedOutputBuffer.Bytes(), err
				},
			}
		}(),
		func() compressionTestCase {
			return compressionTestCase{
				Desc: "ZStandard",
				CompressionCallback: func(rawBytes []byte) ([]byte, error) {
					encoder, err := zstd.NewWriter(nil)
					if err != nil {
						panic(err)
					}
					defer encoder.Close()

					compressedBytes := encoder.EncodeAll(rawBytes, nil)

					return compressedBytes, nil
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					decoder, err := zstd.NewReader(nil)
					if err != nil {
						panic(err)
					}
					defer decoder.Close()

					decompressedBytes, err := decoder.DecodeAll(compressedBytes, nil)
					if err != nil {
						panic(err)
					}

					return decompressedBytes, err
				},
			}
		}(),
		func() compressionTestCase {
			return compressionTestCase{
				Desc: "S2",
				CompressionCallback: func(rawBytes []byte) ([]byte, error) {
					compressedOutputBuffer := &bytes.Buffer{}
					uncompressedInputBuffer := bytes.NewReader(rawBytes)

					compressor := s2.NewWriter(compressedOutputBuffer)
					defer compressor.Close()

					_, err := io.Copy(compressor, uncompressedInputBuffer)
					if err != nil {
						return nil, err
					}

					return compressedOutputBuffer.Bytes(), err
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					compressedBytesReader := bytes.NewReader(compressedBytes)
					decompressedBytesBufferWriter := &bytes.Buffer{}

					decoder := s2.NewReader(compressedBytesReader)

					_, err := io.Copy(decompressedBytesBufferWriter, decoder)
					if err != nil {
						return nil, err
					}

					return decompressedBytesBufferWriter.Bytes(), err
				},
			}
		}(),
		func() compressionTestCase {
			return compressionTestCase{
				Desc: "Deflate",
				CompressionCallback: func(uncompressedRawBytes []byte) ([]byte, error) {
					decompressedBytesBuffer := &bytes.Buffer{}
					uncompressedRawBytesBuffer := bytes.NewReader(uncompressedRawBytes)

					encoder, err := flate.NewWriter(decompressedBytesBuffer, -1)
					if err != nil {
						return nil, err
					}
					defer encoder.Close()

					_, err = io.Copy(encoder, uncompressedRawBytesBuffer)
					if err != nil {
						return nil, err
					}

					return decompressedBytesBuffer.Bytes(), err
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					compressedBytesBufferedReader := bytes.NewReader(compressedBytes)
					decompressedResultBufferedWriter := &bytes.Buffer{}

					decoder := flate.NewReader(compressedBytesBufferedReader)
					defer decoder.Close()

					_, err := io.Copy(decompressedResultBufferedWriter, decoder)
					if err != nil {
						return nil, err
					}

					return decompressedResultBufferedWriter.Bytes(), err
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
