package serialization_deserialization_performance

import (
	"context"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"
)

func Test___SerializationDeserializationWithCompressionPerformance___ThriftCompact(t *testing.T) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftCompactSerializer := arena.NewThriftCompactSerializer()     //binary serializer
	thriftCompactDeserializer := arena.NewThriftCompactDeserializer() //binary deserializer

	for _, test := range arena.AllCompressionTestCases {
		t.Run(test.Desc, func(testbed *testing.T) {

			startTime := time.Now()
			for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
				x := datasource[i%datasourceArrayLength]

				thriftBytes, err := thriftCompactSerializer.Write(ctx, x)
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
				err = thriftCompactDeserializer.Read(ctx, fooitem, decompressedSerializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}
			}
			finishTime := time.Now()

			averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

			testbed.Logf("** ThriftCompact %d nanoseconds\n", int64(averageElapsedTime))
		})
	}
}
