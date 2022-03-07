package serialization_with_compression_performance

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationAndCompressionPerformance___Json(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				jsonBytes, err := json.Marshal(x)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				test.CompressionCallback(jsonBytes)
			}
		})
	}
}
