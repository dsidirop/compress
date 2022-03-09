package serialization_deserialization_performance

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
)

func Test___SerializationDeserializationWithCompressionPerformance___Json(rootTestbed *testing.T) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		rootTestbed.Run(test.Desc, func(testbed *testing.T) {

			startTime := time.Now()
			for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
				x := datasource[i%datasourceArrayLength]

				jsonBytes, err := json.Marshal(x)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(jsonBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				serializedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				fooitem := arena.FooItem{}
				err = json.Unmarshal(serializedBytes, &fooitem)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}
			}
			finishTime := time.Now()

			averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

			fmt.Printf("** JSON+%s %d nanoseconds (avg)\n", test.Desc, int64(averageElapsedTime))
		})
	}
}
