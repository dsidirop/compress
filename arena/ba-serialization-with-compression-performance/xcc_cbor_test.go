package serialization_with_compression_performance

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationAndCompressionPerformance___Cbor(b *testing.B) { // https://github.com/fxamacker/cbor
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				rawBytes, err := cbor.Marshal(x)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				test.CompressionCallback(rawBytes)
			}
		})
	}
}
