package serialization_with_compression_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationWithCompressionMessageSizeFootprint___HambaAvro(rootTestbed *testing.T) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		rootTestbed.Run(test.Desc, func(testbed *testing.T) {

			sumOfEventualBytes := 0
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				rawBytes, err := avro.Marshal(arena.Schemas.GoHambaAvro, &x)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(rawBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				sumOfEventualBytes += len(compressedAndSerializedBytes)
			}

			fmt.Printf("** HambaAvro+%s %d bytes (avg)\n", test.Desc, sumOfEventualBytes/datasourceArrayLength)
		})
	}
}
