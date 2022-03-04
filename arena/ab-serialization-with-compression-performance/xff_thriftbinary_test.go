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
	thriftBinarySerializer := thrift.NewTSerializer() //binary serializer

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer()
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				rawBytes, err := thriftBinarySerializer.Write(ctx, x)
				if err != nil {
					panic(err)
				}

				test.CompressionCallback(rawBytes)
			}
		})
	}
}
