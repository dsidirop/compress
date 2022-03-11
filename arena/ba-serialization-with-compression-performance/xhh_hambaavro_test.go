package serialization_with_compression_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationAndCompressionPerformance___HambaAvro(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				serializedBytes, err := avro.Marshal(x.HambaAvroSchema, x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				test.CompressionCallback(serializedBytes)
			}
		})
	}
}
