package serialization_with_compression_eventual_message_size_footprint

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Test___SerializationWithCompressionMessageSizeFootprint___Json(rootTestbed *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		rootTestbed.Run(test.Desc, func(testbed *testing.T) {

			sumOfEventualBytes := 0
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				jsonBytes, err := json.Marshal(x.Item)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(jsonBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				sumOfEventualBytes += len(compressedAndSerializedBytes)
			}

			fmt.Printf("** JSON+%s %d bytes (avg)\n", test.Desc, sumOfEventualBytes)
		})
	}
}
