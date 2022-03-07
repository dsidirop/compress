package serialization_with_compression_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Benchmark___SerializationAndCompressionPerformance___Protobuf(b *testing.B) {
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Protobuf
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				rawBytes, err := proto.Marshal(x)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				test.CompressionCallback(rawBytes)
			}
		})
	}
}
