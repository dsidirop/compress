package serialization_with_compression_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

func Benchmark___SerializationAndCompressionPerformance___MessagePack(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				jsonBytes, err := msgpack.Marshal(x)
				if err != nil {
					panic(err)
				}

				test.CompressionCallback(jsonBytes)
			}
		})
	}
}
