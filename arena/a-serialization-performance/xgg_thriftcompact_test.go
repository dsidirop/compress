package serialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationPerformance___ThriftCompact(t *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftBinarySerializer := arena.NewThriftCompactSerializer() //compact serializer

	for i := 0; i < t.N; i++ {
		x := datasource[i%datasourceArrayLength]

		_, err := thriftBinarySerializer.Write(ctx, x)
		if err != nil {
			panic(err)
		}
	}
}
