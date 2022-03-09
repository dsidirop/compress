package serialization_deserialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"
)

func Benchmark___SerializationDeserializationWithCompressionPerformance___ThriftCompact(rootBench *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftCompactSerializer := arena.NewThriftCompactSerializer()     //binary serializer
	thriftCompactDeserializer := arena.NewThriftCompactDeserializer() //binary deserializer

	for _, test := range arena.AllCompressionTestCases {
		rootBench.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer() //vital

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				thriftBytes, err := thriftCompactSerializer.Write(ctx, x)
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

				y := thfooitem.NewTHFooItem()
				err = thriftCompactDeserializer.Read(ctx, y, decompressedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
