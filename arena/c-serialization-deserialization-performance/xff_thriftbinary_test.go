package serialization_deserialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"

	"github.com/apache/thrift/lib/go/thrift"
)

func Benchmark___SerializationDeserializationPerformance___ThriftBinary(t *testing.B) {
	y := thfooitem.NewTHFooItem()
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftBinarySerializer := thrift.NewTSerializer()     //binary serializer
	thriftBinaryDeserializer := thrift.NewTDeserializer() //binary deserializer

	for i := 0; i < t.N; i++ {
		x := datasource[i%datasourceArrayLength]

		thriftBytes, err := thriftBinarySerializer.Write(ctx, x)
		if err != nil {
			panic(err)
		}

		err = thriftBinaryDeserializer.Read(ctx, y, thriftBytes)
		if err != nil {
			panic(err)
		}
	}
}
