package decompression_with_deserialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___DecompressionAndDeserializationPerformance___ThriftCompact(b *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			thriftCompactSerializer := arena.NewThriftCompactSerializer()     //compact serializer
			thriftCompactDeserializer := arena.NewThriftCompactDeserializer() //compact deserializer

			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				serializedBytes, err := thriftCompactSerializer.Write(ctx, x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				x.Bytes = compressedAndSerializedBytes
			}

			bench.ResetTimer() //vital
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength] //and now we deserialize and decompress

				decompressedSerializedBytes, err := test.DecompressionCallback(x.Bytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				emptyitem := x.NewEmptyThriftItem()
				err = thriftCompactDeserializer.Read(ctx, emptyitem, decompressedSerializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
