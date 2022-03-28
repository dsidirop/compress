package serialization_deserialization_performance

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
)

var thriftcompactNewItem interface{}

func Benchmark___SerializationDeserializationPerformance___ThriftCompact(b *testing.B) {
	ctx := context.TODO()
	newitem := thrift.TStruct(nil)
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

		newitem = x.NewEmptyThriftItem()
		err = thriftCompactDeserializer.Read(ctx, newitem, thriftBytes)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	thriftcompactNewItem = newitem
}
