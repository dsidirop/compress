package serialization_deserialization_performance

import (
	"fmt"
	"testing"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationDeserializationWithCompressionPerformance___Cbor(rootTestbed *testing.T) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		rootTestbed.Run(test.Desc, func(testbed *testing.T) {

			startTime := time.Now()
			for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
				x := datasource[i%datasourceArrayLength]

				serializedBytes, err := cbor.Marshal(x)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				decompressedSerializedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				fooitem := &arena.FooItem{}
				err = cbor.Unmarshal(decompressedSerializedBytes, fooitem)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}
			}
			finishTime := time.Now()

			averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

			fmt.Printf("** CBOR+%s %d nanoseconds (avg)\n", test.Desc, int64(averageElapsedTime))
		})
	}
}
