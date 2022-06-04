package serialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationPerformance___ThriftCompact(b *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftCompactSerializer := arena.NewThriftCompactSerializer() //compact serializer

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		_, err := thriftCompactSerializer.Write(ctx, x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
