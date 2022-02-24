package serialization_deserialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"
)

func Benchmark___SerializationDeserializationPerformance___ThriftCompact(t *testing.B) {
	y := thfooitem.NewTHFooItem()
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftCompactSerializer := arena.NewThriftCompactSerializer()     // compact serializer
	thriftCompactDeserializer := arena.NewThriftCompactDeserializer() // compact deserializer

	for i := 0; i < t.N; i++ {
		x := datasource[i%datasourceArrayLength]

		thriftBytes, err := thriftCompactSerializer.Write(ctx, x)
		if err != nil {
			panic(err)
		}

		err = thriftCompactDeserializer.Read(ctx, y, thriftBytes)
		if err != nil {
			panic(err)
		}
	}
}
