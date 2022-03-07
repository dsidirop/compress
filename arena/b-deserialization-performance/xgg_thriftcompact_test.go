package deserialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"
)

func Benchmark___DeserializationPerformance___ThriftCompact(b *testing.B) {
	y := thfooitem.NewTHFooItem()
	ctx := context.TODO()
	datasource := arena.SerializedDataSources.ThriftCompact
	datasourceArrayLength := len(datasource)
	thriftBinaryDeserializer := arena.NewThriftCompactDeserializer() //compact deserializer

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		err := thriftBinaryDeserializer.Read(ctx, y, x)
		if err != nil {
			panic(err)
		}
	}
}
