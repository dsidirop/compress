package serialization_deserialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"

	"github.com/apache/thrift/lib/go/thrift"
)

func Benchmark___SerializationDeserializationWithCompressionPerformance___ThriftBinary(rootBench *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftBinarySerializer := thrift.NewTSerializer()     //binary serializer
	thriftBinaryDeserializer := thrift.NewTDeserializer() //binary deserializer

	for _, test := range arena.AllCompressionTestCases {
		rootBench.Run(test.Desc, func(subbench *testing.B) {
			subbench.ResetTimer() //vital

			for i := 0; i < subbench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				thriftBytes, err := thriftBinarySerializer.Write(ctx, x)
				if err != nil {
					subbench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(thriftBytes)
				if err != nil {
					subbench.Fatalf("Error: %s", err)
				}

				decompressedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					subbench.Fatalf("Error: %s", err)
				}

				y := thfooitem.NewTHFooItem()
				err = thriftBinaryDeserializer.Read(ctx, y, decompressedBytes)
				if err != nil {
					subbench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
