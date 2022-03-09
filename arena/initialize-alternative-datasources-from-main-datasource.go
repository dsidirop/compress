package arena

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/andybalholm/brotli"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/fxamacker/cbor/v2"
	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena/thfooitem"
	"github.com/klauspost/compress/flate"
	"github.com/klauspost/compress/s2"
	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4/v4"
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
					compressedBytesBufferWriter := &bytes.Buffer{}

					zlibCompressor := zlib.NewWriter(compressedBytesBufferWriter)

					_, err := zlibCompressor.Write(rawBytes)
					zlibCompressor.Close() //dont use defer   it wont work

					if err != nil {
						return nil, err
					}

					return compressedBytesBufferWriter.Bytes(), err
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					compressedInputBuffer := bytes.NewReader(compressedBytes)

					zlibDecompressor, err := zlib.NewReader(compressedInputBuffer)
					if err != nil {
						return nil, err
					}

					decompressedBytes, err := ioutil.ReadAll(zlibDecompressor)
					zlibDecompressor.Close()

					if err != nil {
						return nil, err
					}

					return decompressedBytes, nil
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

					compressedBytes := encoder.EncodeAll(rawBytes, nil)
					encoder.Close()

					return compressedBytes, nil
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					decoder, err := zstd.NewReader(nil)
					if err != nil {
						panic(err)
					}

					decompressedBytes, err := decoder.DecodeAll(compressedBytes, nil)
					decoder.Close()

					if err != nil {
						panic(err)
					}

					return decompressedBytes, err
				},
			}
		}(),
		func() compressionTestCase {
			return compressionTestCase{
				Desc: "Snappy2", //via s2 which means snappy2   its a facelift over the original snappy
				CompressionCallback: func(rawBytes []byte) ([]byte, error) {
					compressedBytes := s2.Encode(nil, rawBytes)

					return compressedBytes, nil
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					decompressedBytes, err := s2.Decode(nil, compressedBytes)
					if err != nil {
						return nil, err
					}

					return decompressedBytes, nil
				},
			}
		}(),
		func() compressionTestCase {
			return compressionTestCase{
				Desc: "Deflate",
				CompressionCallback: func(rawBytes []byte) ([]byte, error) {
					decompressedBytesBuffer := &bytes.Buffer{}
					uncompressedRawBytesBuffer := bytes.NewReader(rawBytes)

					encoder, err := flate.NewWriter(decompressedBytesBuffer, -1)
					if err != nil {
						return nil, err
					}

					_, err = io.Copy(encoder, uncompressedRawBytesBuffer)
					encoder.Close()

					if err != nil {
						return nil, err
					}

					decompressedBytes := decompressedBytesBuffer.Bytes()

					return decompressedBytes, err
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					compressedBytesBufferedReader := bytes.NewReader(compressedBytes)

					flateDecompressor := flate.NewReader(compressedBytesBufferedReader)

					decompressedBytes, err := ioutil.ReadAll(flateDecompressor)
					flateDecompressor.Close()

					if err != nil {
						return nil, err
					}

					return decompressedBytes, nil
				},
			}
		}(),
		func() compressionTestCase {
			return compressionTestCase{
				Desc: "GZip",
				CompressionCallback: func(rawBytes []byte) ([]byte, error) {
					compressedBytesBuffer := &bytes.Buffer{}
					gzipCompressor := gzip.NewWriter(compressedBytesBuffer)

					_, err := gzipCompressor.Write(rawBytes)
					if err != nil {
						return nil, err
					}

					err = gzipCompressor.Close()
					if err != nil {
						return nil, err
					}

					return compressedBytesBuffer.Bytes(), nil
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					compressedBytesReader := bytes.NewReader(compressedBytes)

					gzipDecompressor, err := gzip.NewReader(compressedBytesReader)
					if err != nil {
						return nil, err
					}

					decompressedBytes, err := ioutil.ReadAll(gzipDecompressor)
					gzipDecompressor.Close()

					if err != nil {
						return nil, err
					}

					return decompressedBytes, nil
				},
			}
		}(),
		func() compressionTestCase {
			return compressionTestCase{
				Desc: "Brotli",
				CompressionCallback: func(rawBytes []byte) ([]byte, error) {
					compressedOutputBuffer := &bytes.Buffer{}

					brotliCompressor := brotli.NewWriter(compressedOutputBuffer)

					_, err := brotliCompressor.Write(rawBytes)
					if err != nil {
						return nil, err
					}

					err = brotliCompressor.Close()
					if err != nil {
						return nil, err
					}

					return compressedOutputBuffer.Bytes(), nil
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					compressedBytesBuffer := bytes.NewReader(compressedBytes)

					brotliDecompressor := brotli.NewReader(compressedBytesBuffer)

					decompressedBytes, err := ioutil.ReadAll(brotliDecompressor)
					//brotliDecompressor.Close()

					if err != nil {
						return nil, err
					}

					return decompressedBytes, nil
				},
			}
		}(),
		func() compressionTestCase {
			return compressionTestCase{
				Desc: "LZ4",
				CompressionCallback: func(rawBytes []byte) ([]byte, error) {
					compressedOutputBuffer := &bytes.Buffer{}
					lz4Compressor := lz4.NewWriter(compressedOutputBuffer)

					_, err := lz4Compressor.Write(rawBytes)
					lz4Compressor.Close()

					if err != nil {
						return nil, err
					}

					return compressedOutputBuffer.Bytes(), nil
				},
				DecompressionCallback: func(compressedBytes []byte) ([]byte, error) {
					compressedBytesBuffer := bytes.NewReader(compressedBytes)

					lz4Decompressor := lz4.NewReader(compressedBytesBuffer)

					decompressedBytes, err := ioutil.ReadAll(lz4Decompressor)
					//lz4Decompressor.Close()

					if err != nil {
						return nil, err
					}

					return decompressedBytes, nil
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
