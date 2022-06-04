package serialization_with_compression_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationAndCompressionPerformance___ThriftCompact(b *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftCompactSerializer := arena.NewThriftCompactSerializer() //compact serializer

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				rawBytes, err := thriftCompactSerializer.Write(ctx, x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				test.CompressionCallback(rawBytes)
			}
		})
	}
}
