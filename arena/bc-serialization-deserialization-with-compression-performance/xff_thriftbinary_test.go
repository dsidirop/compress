package serialization_deserialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"

	"github.com/apache/thrift/lib/go/thrift"
)

func Benchmark___SerializationDeserializationPerformance___ThriftBinary(b *testing.B) {
	y := thfooitem.NewTHFooItem()
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftBinarySerializer := thrift.NewTSerializer()     //binary serializer
	thriftBinaryDeserializer := thrift.NewTDeserializer() //binary deserializer

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		thriftBytes, err := thriftBinarySerializer.Write(ctx, x)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		err = thriftBinaryDeserializer.Read(ctx, y, thriftBytes)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
