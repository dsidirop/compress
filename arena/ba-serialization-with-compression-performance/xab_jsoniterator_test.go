package serialization_with_compression_performance

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/klauspost/compress/arena"
)

var jsonIteratorBytes []byte

func Benchmark___SerializationAndCompressionPerformance___JsonIterator(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	var results []byte
	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				jsonBytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				results, err = test.CompressionCallback(jsonBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})

		jsonIteratorBytes = results
	}
}
