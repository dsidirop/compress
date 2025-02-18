package serialization_with_compression_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

var protobufCompressedAndSerializedBytes []byte //keep global

func Test___SerializationWithCompressionMessageSizeFootprint___Protobuf(rootTestbed *testing.T) {
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Protobuf
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		compressedAndSerializedBytes := []byte(nil)

		rootTestbed.Run(test.Desc, func(testbed *testing.T) {
			sumOfEventualBytes := 0
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				rawBytes, err := proto.Marshal(x.Item)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(rawBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				sumOfEventualBytes += len(compressedAndSerializedBytes)

				fmt.Printf("* Protobuf +%d bytes\n", len(compressedAndSerializedBytes))
			}

			fmt.Printf("** Protobuf+%s %d bytes (avg)\n", test.Desc, sumOfEventualBytes)

			protobufCompressedAndSerializedBytes = compressedAndSerializedBytes
		})
	}
}
