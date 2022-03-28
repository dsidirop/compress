package serialization_with_compression_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/mailru/easyjson"
)

var jsoneasyCompressedAndSerializedBytes []byte //keep global

func Test___SerializationWithCompressionMessageSizeFootprint___JsonEasy(rootTestbed *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		compressedAndSerializedBytes := []byte(nil)

		rootTestbed.Run(test.Desc, func(testbed *testing.T) {
			sumOfEventualBytes := 0
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				jsonBytes, err := easyjson.Marshal(x.Item)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err = test.CompressionCallback(jsonBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				sumOfEventualBytes += len(compressedAndSerializedBytes)

				fmt.Printf("* JSONEasy +%d bytes\n", len(compressedAndSerializedBytes))
			}

			fmt.Printf("** JSONEasy+%s %d bytes (avg)\n", test.Desc, sumOfEventualBytes)
		})

		jsoneasyCompressedAndSerializedBytes = compressedAndSerializedBytes
	}
}
