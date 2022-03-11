package deserialization_performance

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
)

func Benchmark___DeserializationPerformance___ThriftBinary(b *testing.B) {
	ctx := context.TODO()
	datasource := arena.SerializedDataSources.ThriftBinary
	datasourceArrayLength := len(datasource)
	thriftBinaryDeserializer := thrift.NewTDeserializer() //binary deserializer

	b.ResetTimer()
	for iterator := 0; iterator < b.N; iterator++ {
		i := iterator % datasourceArrayLength

		newitem := arena.MainDatasource[i].NewEmptyThriftItem()
		rawbytes := datasource[i]

		err := thriftBinaryDeserializer.Read(ctx, newitem, rawbytes)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
