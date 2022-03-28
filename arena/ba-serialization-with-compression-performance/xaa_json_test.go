package serialization_with_compression_performance

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

var jsonBytes []byte

func Benchmark___SerializationAndCompressionPerformance___Json(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	var results []byte
	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				jsonBytes, err := json.Marshal(x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				results, err = test.CompressionCallback(jsonBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})

		jsonBytes = results
	}
}
