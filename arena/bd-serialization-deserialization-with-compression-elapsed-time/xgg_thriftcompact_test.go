package serialization_deserialization_performance

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
)

var thriftcompactDummyVariable interface{} //keep global

func Test___SerializationDeserializationWithCompressionPerformance___ThriftCompact(rootTestbed *testing.T) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftCompactSerializer := arena.NewThriftCompactSerializer()     //binary serializer
	thriftCompactDeserializer := arena.NewThriftCompactDeserializer() //binary deserializer

	for _, test := range arena.AllCompressionTestCases {
		rootTestbed.Run(test.Desc, func(testbed *testing.T) {

			newitem := thrift.TStruct(nil)

			startTime := time.Now()
			for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
				x := datasource[i%datasourceArrayLength]

				thriftBytes, err := thriftCompactSerializer.Write(ctx, x.Item)
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

				newitem = x.NewEmptyThriftItem()
				err = thriftCompactDeserializer.Read(ctx, newitem, decompressedSerializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}
			}

			finishTime := time.Now()
			averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

			fmt.Printf("** ThriftCompact+%s %d nanoseconds (avg)\n", test.Desc, int64(averageElapsedTime))

			thriftcompactDummyVariable = newitem
		})
	}
}
