package deserialization_performance

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"
)

func Benchmark___DeserializationPerformance___ThriftBinary(t *testing.B) {
	y := thfooitem.NewTHFooItem()
	ctx := context.TODO()
	datasource := arena.SerializedDataSources.ThriftBinary
	datasourceArrayLength := len(datasource)
	thriftBinaryDeserializer := thrift.NewTDeserializer() //binary deserializer

	for i := 0; i < t.N; i++ {
		x := datasource[i%datasourceArrayLength]

		err := thriftBinaryDeserializer.Read(ctx, y, x)
		if err != nil {
			panic(err)
		}
	}
}
