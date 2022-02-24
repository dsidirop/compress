package serialization_performance

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationPerformance___ThriftBinary(t *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftBinarySerializer := thrift.NewTSerializer() //binary serializer

	for i := 0; i < t.N; i++ {
		x := datasource[i%datasourceArrayLength]

		_, err := thriftBinarySerializer.Write(ctx, x)
		if err != nil {
			panic(err)
		}
	}
}
