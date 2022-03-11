package serialization_deserialization_performance

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Test___SerializationDeserializationWithCompressionPerformance___Msgp(rootTestbed *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		rootTestbed.Run(test.Desc, func(testbed *testing.T) {
			startTime := time.Now()

			for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
				x := datasource[i%datasourceArrayLength]

				serializedBytesBuffer := &bytes.Buffer{}
				err := msgp.Encode(serializedBytesBuffer, x.Item)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytesBuffer.Bytes())
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				serializedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				newitem := x.NewEmptyItem()
				err = msgp.Decode(bytes.NewReader(serializedBytes), newitem)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}
			}
			finishTime := time.Now()

			averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

			fmt.Printf("** Msgp+%s %d nanoseconds (avg)\n", test.Desc, int64(averageElapsedTime))
		})
	}
}
