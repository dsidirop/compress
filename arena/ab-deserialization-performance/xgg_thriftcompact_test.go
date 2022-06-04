package deserialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___DeserializationPerformance___ThriftCompact(b *testing.B) {
	ctx := context.TODO()
	datasource := arena.SerializedDataSources.ThriftCompact
	datasourceArrayLength := len(datasource)
	thriftCompactDeserializer := arena.NewThriftCompactDeserializer() //compact deserializer

	b.ResetTimer()
	for iterator := 0; iterator < b.N; iterator++ {
		i := iterator % datasourceArrayLength

		newitem := arena.MainDatasource[i].NewEmptyThriftItem()
		rawbytes := datasource[i]

		err := thriftCompactDeserializer.Read(ctx, newitem, rawbytes)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
