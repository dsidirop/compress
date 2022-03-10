package arena

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/andybalholm/brotli"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/dsnet/compress/bzip2"
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
	//zlib
	zlibCompressorFactory := func(compressionLevel int) func(rawBytes []byte) ([]byte, error) {
		return func(rawBytes []byte) ([]byte, error) {
			compressedBytesBufferWriter := &bytes.Buffer{}

			zlibCompressor, err := zlib.NewWriterLevel(compressedBytesBufferWriter, compressionLevel)
			if err != nil {
				return nil, err
			}

			_, err = zlibCompressor.Write(rawBytes)
			zlibCompressor.Close() //dont use defer   it wont work

			if err != nil {
				return nil, err
			}

			return compressedBytesBufferWriter.Bytes(), err
		}
	}
	zlibDecompressorFactory := func() func(compressedBytes []byte) ([]byte, error) {
		return func(compressedBytes []byte) ([]byte, error) {
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
		}
	}
	//zstandard
	zstandardCompressorFactory := func(compressionLevel zstd.EncoderLevel) func(rawBytes []byte) ([]byte, error) {
		return func(rawBytes []byte) ([]byte, error) {
			encoder, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(compressionLevel))
			if err != nil {
				panic(err)
			}

			compressedBytes := encoder.EncodeAll(rawBytes, nil)
			encoder.Close() //vital

			return compressedBytes, nil
		}
	}
	zstandardDecompressorFactory := func() func(compressedBytes []byte) ([]byte, error) {
		return func(compressedBytes []byte) ([]byte, error) {
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
		}
	}
	//s2
	s2CompressorFactory := func(compressionLevel int) func(rawBytes []byte) ([]byte, error) {
		if compressionLevel == 0 {
			return func(rawBytes []byte) ([]byte, error) {
				compressedBytes := s2.Encode(nil, rawBytes) //default

				return compressedBytes, nil
			}
		}

		if compressionLevel == 1 {
			return func(rawBytes []byte) ([]byte, error) {
				compressedBytes := s2.EncodeBetter(nil, rawBytes)

				return compressedBytes, nil
			}
		}

		if compressionLevel == 2 {
			return func(rawBytes []byte) ([]byte, error) {
				compressedBytes := s2.EncodeBest(nil, rawBytes)

				return compressedBytes, nil
			}
		}

		panic(fmt.Sprintf("s2compressorfactory: unsupported compression level %d", compressionLevel))
	}
	s2DecompressorFactory := func() func(compressedBytes []byte) ([]byte, error) {
		return func(compressedBytes []byte) ([]byte, error) {
			decompressedBytes, err := s2.Decode(nil, compressedBytes)
			if err != nil {
				return nil, err
			}

			return decompressedBytes, nil
		}
	}
	//snappy
	snappyCompressorFactory := func(compressionLevel int) func(rawBytes []byte) ([]byte, error) {
		if compressionLevel == 0 {
			return func(rawBytes []byte) ([]byte, error) {
				compressedBytes := s2.EncodeSnappy(nil, rawBytes) //default

				return compressedBytes, nil
			}
		}

		if compressionLevel == 1 {
			return func(rawBytes []byte) ([]byte, error) {
				compressedBytes := s2.EncodeSnappyBetter(nil, rawBytes)

				return compressedBytes, nil
			}
		}

		if compressionLevel == 2 {
			return func(rawBytes []byte) ([]byte, error) {
				compressedBytes := s2.EncodeSnappyBest(nil, rawBytes)

				return compressedBytes, nil
			}
		}

		panic(fmt.Sprintf("snappycompressorfactory: unsupported compression level %d", compressionLevel))
	}
	snappyDecompressorFactory := func() func(compressedBytes []byte) ([]byte, error) {
		return func(compressedBytes []byte) ([]byte, error) {
			decompressedBytes, err := s2.Decode(nil, compressedBytes)
			if err != nil {
				return nil, err
			}

			return decompressedBytes, nil
		}
	}
	//deflate
	deflateCompressorFactory := func(compressionLevel int) func(compressedBytes []byte) ([]byte, error) {
		return func(rawBytes []byte) ([]byte, error) {
			decompressedBytesBuffer := &bytes.Buffer{}
			uncompressedRawBytesBuffer := bytes.NewReader(rawBytes)

			encoder, err := flate.NewWriter(decompressedBytesBuffer, compressionLevel)
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
		}
	}
	deflateDecompressorFactory := func() func(compressedBytes []byte) ([]byte, error) {
		return func(compressedBytes []byte) ([]byte, error) {
			compressedBytesBufferedReader := bytes.NewReader(compressedBytes)

			flateDecompressor := flate.NewReader(compressedBytesBufferedReader)

			decompressedBytes, err := ioutil.ReadAll(flateDecompressor)
			flateDecompressor.Close()

			if err != nil {
				return nil, err
			}

			return decompressedBytes, nil
		}
	}
	//gzip
	gzipCompressorFactory := func(compressionLevel int) func(compressedBytes []byte) ([]byte, error) {
		return func(rawBytes []byte) ([]byte, error) {
			compressedBytesBuffer := &bytes.Buffer{}
			gzipCompressor, err := gzip.NewWriterLevel(compressedBytesBuffer, compressionLevel)
			if err != nil {
				return nil, err
			}

			_, err = gzipCompressor.Write(rawBytes)
			if err != nil {
				return nil, err
			}

			err = gzipCompressor.Close()
			if err != nil {
				return nil, err
			}

			return compressedBytesBuffer.Bytes(), nil
		}
	}
	gzipDecompressorFactory := func() func(compressedBytes []byte) ([]byte, error) {
		return func(compressedBytes []byte) ([]byte, error) {
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
		}
	}
	//brotli
	brotliCompressorFactory := func(compressionLevel int) func(compressedBytes []byte) ([]byte, error) {
		return func(rawBytes []byte) ([]byte, error) {
			compressedOutputBuffer := &bytes.Buffer{}
			brotliCompressor := brotli.NewWriterLevel(compressedOutputBuffer, compressionLevel)

			_, err := brotliCompressor.Write(rawBytes)
			if err != nil {
				return nil, err
			}

			err = brotliCompressor.Close()
			if err != nil {
				return nil, err
			}

			return compressedOutputBuffer.Bytes(), nil
		}
	}
	brotliDecompressorFactory := func() func(compressedBytes []byte) ([]byte, error) {
		return func(compressedBytes []byte) ([]byte, error) {
			compressedBytesBuffer := bytes.NewReader(compressedBytes)

			brotliDecompressor := brotli.NewReader(compressedBytesBuffer)

			decompressedBytes, err := ioutil.ReadAll(brotliDecompressor)
			//brotliDecompressor.Close()

			if err != nil {
				return nil, err
			}

			return decompressedBytes, nil
		}
	}
	//lz4
	lz4CompressorFactory := func(compressionLevel *lz4.CompressionLevel) func(compressedBytes []byte) ([]byte, error) {
		return func(rawBytes []byte) ([]byte, error) {
			compressedOutputBuffer := &bytes.Buffer{}
			lz4Compressor := lz4.NewWriter(compressedOutputBuffer) //lz4Compressor.Reset() todo   experiment with this

			if compressionLevel != nil {
				lz4Compressor.Apply(lz4.CompressionLevelOption(*compressionLevel))
			}

			_, err := lz4Compressor.Write(rawBytes)
			lz4Compressor.Close()

			if err != nil {
				return nil, err
			}

			return compressedOutputBuffer.Bytes(), nil
		}
	}
	lz4DecompressorFactory := func() func(compressedBytes []byte) ([]byte, error) {
		return func(compressedBytes []byte) ([]byte, error) {
			compressedBytesBuffer := bytes.NewReader(compressedBytes)

			lz4Decompressor := lz4.NewReader(compressedBytesBuffer)

			decompressedBytes, err := ioutil.ReadAll(lz4Decompressor)
			//lz4Decompressor.Close()

			if err != nil {
				return nil, err
			}

			return decompressedBytes, nil
		}
	}
	//bzip2
	bzip2CompressorFactory := func(compressionLevel int) func(compressedBytes []byte) ([]byte, error) {
		return func(rawBytes []byte) ([]byte, error) {
			compressedOutputBuffer := &bytes.Buffer{}

			bzip2Compressor, err := bzip2.NewWriter(compressedOutputBuffer, &bzip2.WriterConfig{})
			if err != nil {
				return nil, err
			}

			_, err = bzip2Compressor.Write(rawBytes)
			bzip2Compressor.Close()

			if err != nil {
				return nil, err
			}

			return compressedOutputBuffer.Bytes(), nil
		}
	}
	bzip2DecompressorFactory := func() func(compressedBytes []byte) ([]byte, error) {
		return func(compressedBytes []byte) ([]byte, error) {
			compressedBytesBuffer := bytes.NewReader(compressedBytes)

			bzip2Decompressor, err := bzip2.NewReader(compressedBytesBuffer, &bzip2.ReaderConfig{})
			if err != nil {
				return nil, err
			}

			decompressedBytes, err := ioutil.ReadAll(bzip2Decompressor)
			bzip2Decompressor.Close()

			if err != nil {
				return nil, err
			}

			return decompressedBytes, nil
		}
	}

	AllCompressionTestCases = []compressionTestCase{
		//zlib
		{
			Desc:                  "Zlib-DefaultCompression",
			CompressionCallback:   zlibCompressorFactory(zlib.DefaultCompression),
			DecompressionCallback: zlibDecompressorFactory(),
		},
		{
			Desc:                  "Zlib-BestCompression",
			CompressionCallback:   zlibCompressorFactory(zlib.BestCompression),
			DecompressionCallback: zlibDecompressorFactory(),
		},
		{
			Desc:                  "Zlib-BestSpeed",
			CompressionCallback:   zlibCompressorFactory(zlib.BestSpeed),
			DecompressionCallback: zlibDecompressorFactory(),
		},
		//zstandard
		{
			Desc:                  "ZStandard-DefaultCompression",
			CompressionCallback:   zstandardCompressorFactory(zstd.SpeedDefault),
			DecompressionCallback: zstandardDecompressorFactory(),
		},
		{
			Desc:                  "ZStandard-BetterCompression", //best compression is ultra dead-slow
			CompressionCallback:   zstandardCompressorFactory(zstd.SpeedBetterCompression),
			DecompressionCallback: zstandardDecompressorFactory(),
		},
		{
			Desc:                  "ZStandard-BestSpeed",
			CompressionCallback:   zstandardCompressorFactory(zstd.SpeedFastest),
			DecompressionCallback: zstandardDecompressorFactory(),
		},
		//s2
		{
			Desc:                  "S2-DefaultCompression",
			CompressionCallback:   s2CompressorFactory(0),
			DecompressionCallback: s2DecompressorFactory(),
		},
		{
			Desc:                  "S2-BetterCompression",
			CompressionCallback:   s2CompressorFactory(1),
			DecompressionCallback: s2DecompressorFactory(),
		},
		{
			Desc:                  "S2-BestCompression",
			CompressionCallback:   s2CompressorFactory(2),
			DecompressionCallback: s2DecompressorFactory(),
		},
		//snappy
		{
			Desc:                  "Snappy-DefaultCompression",
			CompressionCallback:   snappyCompressorFactory(0),
			DecompressionCallback: snappyDecompressorFactory(),
		},
		{
			Desc:                  "Snappy-BetterCompression",
			CompressionCallback:   snappyCompressorFactory(1),
			DecompressionCallback: snappyDecompressorFactory(),
		},
		{
			Desc:                  "Snappy-BestCompression",
			CompressionCallback:   snappyCompressorFactory(2),
			DecompressionCallback: snappyDecompressorFactory(),
		},
		//deflate
		{
			Desc:                  "Deflate-DefaultCompression",
			CompressionCallback:   deflateCompressorFactory(flate.DefaultCompression),
			DecompressionCallback: deflateDecompressorFactory(),
		},
		{
			Desc:                  "Deflate-BestSpeed",
			CompressionCallback:   deflateCompressorFactory(flate.BestSpeed),
			DecompressionCallback: deflateDecompressorFactory(),
		},
		{
			Desc:                  "Deflate-BestCompression",
			CompressionCallback:   deflateCompressorFactory(flate.BestCompression),
			DecompressionCallback: deflateDecompressorFactory(),
		},
		//gzip
		{
			Desc:                  "Gzip-DefaultCompression",
			CompressionCallback:   gzipCompressorFactory(flate.DefaultCompression),
			DecompressionCallback: gzipDecompressorFactory(),
		},
		{
			Desc:                  "Gzip-BestSpeed",
			CompressionCallback:   gzipCompressorFactory(flate.BestSpeed),
			DecompressionCallback: gzipDecompressorFactory(),
		},
		{
			Desc:                  "Gzip-BestCompression",
			CompressionCallback:   gzipCompressorFactory(flate.BestCompression),
			DecompressionCallback: gzipDecompressorFactory(),
		},
		//brotli
		{
			Desc:                  "Brotli-DefaultCompression",
			CompressionCallback:   brotliCompressorFactory(brotli.DefaultCompression),
			DecompressionCallback: brotliDecompressorFactory(),
		},
		{
			Desc:                  "Brotli-BestSpeed",
			CompressionCallback:   brotliCompressorFactory(brotli.BestSpeed),
			DecompressionCallback: brotliDecompressorFactory(),
		},
		{
			Desc:                  "Brotli-BestCompression",
			CompressionCallback:   brotliCompressorFactory(brotli.BestCompression),
			DecompressionCallback: brotliDecompressorFactory(),
		},
		//lz4
		{
			Desc:                  "LZ4-DefaultCompression",
			CompressionCallback:   lz4CompressorFactory(nil),
			DecompressionCallback: lz4DecompressorFactory(),
		},
		{
			Desc:                  "LZ4-BestSpeed",
			CompressionCallback:   lz4CompressorFactory(&[]lz4.CompressionLevel{lz4.Fast}[0]),
			DecompressionCallback: lz4DecompressorFactory(),
		},
		{
			Desc:                  "LZ4-BestCompression",
			CompressionCallback:   lz4CompressorFactory(&[]lz4.CompressionLevel{lz4.Level9}[0]),
			DecompressionCallback: lz4DecompressorFactory(),
		},
		//bzip2
		{
			Desc:                  "Bzip2-DefaultCompression",
			CompressionCallback:   bzip2CompressorFactory(bzip2.DefaultCompression),
			DecompressionCallback: bzip2DecompressorFactory(),
		},
		{
			Desc:                  "Bzip2-BestSpeed",
			CompressionCallback:   bzip2CompressorFactory(bzip2.BestSpeed),
			DecompressionCallback: bzip2DecompressorFactory(),
		},
		{
			Desc:                  "Bzip2-BestCompression",
			CompressionCallback:   bzip2CompressorFactory(bzip2.BestCompression),
			DecompressionCallback: bzip2DecompressorFactory(),
		},
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
