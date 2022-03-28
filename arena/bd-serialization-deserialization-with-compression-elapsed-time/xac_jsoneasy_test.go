package serialization_deserialization_performance

import (
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"github.com/mailru/easyjson"
)

var jsoneasyIteratorDummyVariable interface{} //keep global

func Test___SerializationDeserializationWithCompressionPerformance___JsonEasy(rootTestbed *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		rootTestbed.Run(test.Desc, func(testbed *testing.T) {

			newitem := arena.ItemSerdeable(nil)

			startTime := time.Now()
			for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
				x := datasource[i%datasourceArrayLength]

				jsonBytes, err := easyjson.Marshal(x.Item)
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

				newitem = x.NewEmptyItem()
				err = easyjson.Unmarshal(serializedBytes, newitem)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}
			}

			finishTime := time.Now()

			averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

			fmt.Printf("** JSONEasy+%s %d nanoseconds (avg)\n", test.Desc, int64(averageElapsedTime))

			jsoneasyIteratorDummyVariable = newitem
		})
	}
}
