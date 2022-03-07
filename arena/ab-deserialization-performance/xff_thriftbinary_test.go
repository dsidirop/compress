package deserialization_performance

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"
)

func Benchmark___DeserializationPerformance___ThriftBinary(b *testing.B) {
	y := thfooitem.NewTHFooItem()
	ctx := context.TODO()
	datasource := arena.SerializedDataSources.ThriftBinary
	datasourceArrayLength := len(datasource)
	thriftBinaryDeserializer := thrift.NewTDeserializer() //binary deserializer

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		err := thriftBinaryDeserializer.Read(ctx, y, x)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
