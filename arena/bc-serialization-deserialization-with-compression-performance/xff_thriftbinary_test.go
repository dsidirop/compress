package serialization_deserialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"

	"github.com/apache/thrift/lib/go/thrift"
)

var thriftbinaryDummyVariable interface{}

func Benchmark___SerializationDeserializationWithCompressionPerformance___ThriftBinary(rootBench *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftBinarySerializer := thrift.NewTSerializer()     //binary serializer
	thriftBinaryDeserializer := thrift.NewTDeserializer() //binary deserializer

	for _, test := range arena.AllCompressionTestCases {
		rootBench.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer() //vital

			newitem := thrift.TStruct(nil)
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				thriftBytes, err := thriftBinarySerializer.Write(ctx, x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(thriftBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				decompressedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				newitem = x.NewEmptyThriftItem()
				err = thriftBinaryDeserializer.Read(ctx, newitem, decompressedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}

			thriftbinaryDummyVariable = newitem
		})
	}
}
