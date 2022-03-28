package serialization_with_compression_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/mailru/easyjson"
)

var jsoneasyIteratorBytes []byte

func Benchmark___SerializationAndCompressionPerformance___JsonEasy(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	var results []byte
	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				jsonBytes, err := easyjson.Marshal(x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				results, err = test.CompressionCallback(jsonBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})

		jsoneasyIteratorBytes = results
	}
}
