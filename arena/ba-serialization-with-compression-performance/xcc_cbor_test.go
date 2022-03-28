package serialization_with_compression_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
)

var cborBytes []byte

func Benchmark___SerializationAndCompressionPerformance___Cbor(b *testing.B) { // https://github.com/fxamacker/cbor
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	var results []byte
	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				rawBytes, err := arena.CborStandardEncoder.Marshal(x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				results, err = test.CompressionCallback(rawBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})

		cborBytes = results
	}
}
