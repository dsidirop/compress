package serialization_deserialization_elapsed_time

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"
)

func Test___SerializationDeserializationElapsedTime___ThriftBinary(t *testing.T) {
	y := thfooitem.NewTHFooItem()
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftBinarySerializer := thrift.NewTSerializer()     //binary serializer
	thriftBinaryDeserializer := thrift.NewTDeserializer() //binary deserializer

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
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
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** ThriftBinary %d nanoseconds\n", int64(averageElapsedTime))
}
