package serialization_with_compression_performance

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationAndCompressionPerformance___ThriftBinary(b *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				thriftBinarySerializer := thrift.NewTSerializer() //binary serializer

				rawBytes, err := thriftBinarySerializer.Write(ctx, x)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				// if test.Desc == "S2" {
				// 	continue
				// }

				test.CompressionCallback(rawBytes)
			}
		})
	}
}
