package serialization_with_compression_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
)

var thriftcompactBytes []byte

func Benchmark___SerializationAndCompressionPerformance___ThriftCompact(b *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftCompactSerializer := arena.NewThriftCompactSerializer() //compact serializer

	var results []byte
	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				rawBytes, err := thriftCompactSerializer.Write(ctx, x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				results, err = test.CompressionCallback(rawBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})

		thriftcompactBytes = results
	}
}
