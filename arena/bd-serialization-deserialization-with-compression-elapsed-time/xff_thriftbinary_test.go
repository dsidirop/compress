package serialization_deserialization_performance

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"

	"github.com/apache/thrift/lib/go/thrift"
)

func Test___SerializationDeserializationWithCompressionPerformance___ThriftBinary(rootTestbed *testing.T) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftBinarySerializer := thrift.NewTSerializer()     //binary serializer
	thriftBinaryDeserializer := thrift.NewTDeserializer() //binary deserializer

	for _, test := range arena.AllCompressionTestCases {
		rootTestbed.Run(test.Desc, func(testbed *testing.T) {

			startTime := time.Now()
			for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
				x := datasource[i%datasourceArrayLength]

				thriftBytes, err := thriftBinarySerializer.Write(ctx, x)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(thriftBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				decompressedSerializedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				fooitem := thfooitem.NewTHFooItem()
				err = thriftBinaryDeserializer.Read(ctx, fooitem, decompressedSerializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}
			}
			finishTime := time.Now()

			averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

			fmt.Printf("** ThriftBinary+%s %d nanoseconds\n", test.Desc, int64(averageElapsedTime))
		})
	}
}
