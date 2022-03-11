package serialization_performance

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationPerformance___ThriftBinary(b *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftBinarySerializer := thrift.NewTSerializer() //binary serializer

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		_, err := thriftBinarySerializer.Write(ctx, x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
