package serialization_with_compression_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

func Test___SerializationWithCompressionMessageSizeFootprint___Bson(rootTestbed *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		rootTestbed.Run(test.Desc, func(testbed *testing.T) {

			sumOfEventualBytes := 0
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				rawBytes, err := bson.Marshal(x.Item)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(rawBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				sumOfEventualBytes += len(compressedAndSerializedBytes)
			}

			fmt.Printf("** Bson+%s %d bytes (avg)\n", test.Desc, sumOfEventualBytes/datasourceArrayLength)
		})
	}
}
