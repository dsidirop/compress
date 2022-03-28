package serialization_with_compression_eventual_message_size_footprint

import (
	"context"
	"fmt"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
)

var thriftbinaryCompressedAndSerializedBytes []byte //keep global

func Test___SerializationWithCompressionMessageSizeFootprint___ThriftBinary(rootTestbed *testing.T) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftBinarySerializer := thrift.NewTSerializer() //binary serializer

	for _, test := range arena.AllCompressionTestCases {
		compressedAndSerializedBytes := []byte(nil)

		rootTestbed.Run(test.Desc, func(testbed *testing.T) {
			sumOfEventualBytes := 0
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				rawBytes, err := thriftBinarySerializer.Write(ctx, x.Item)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(rawBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				sumOfEventualBytes += len(compressedAndSerializedBytes)

				fmt.Printf("* ThriftBinary +%d bytes\n", len(compressedAndSerializedBytes))
			}

			fmt.Printf("** ThriftBinary+%s %d bytes (avg)\n", test.Desc, sumOfEventualBytes)

			thriftbinaryCompressedAndSerializedBytes = compressedAndSerializedBytes
		})
	}
}
