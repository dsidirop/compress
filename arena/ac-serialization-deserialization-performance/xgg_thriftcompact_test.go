package serialization_deserialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"
)

func Benchmark___SerializationDeserializationPerformance___ThriftCompact(b *testing.B) {
	ctx := context.TODO()
	fooitem := thfooitem.NewTHFooItem()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftCompactSerializer := arena.NewThriftCompactSerializer()     // compact serializer
	thriftCompactDeserializer := arena.NewThriftCompactDeserializer() // compact deserializer

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		thriftBytes, err := thriftCompactSerializer.Write(ctx, x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		err = thriftCompactDeserializer.Read(ctx, fooitem, thriftBytes)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
