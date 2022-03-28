package serialization_eventual_message_size_footprint

import (
	"context"
	"fmt"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___ThriftBinary(t *testing.T) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftBinarySerializer := thrift.NewTSerializer() //binary serializer

	totalBytesCount := 0
	for i := 0; i < datasourceArrayLength; i++ {
		x := datasource[i]

		rawBytes, err := thriftBinarySerializer.Write(ctx, x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		totalBytesCount += len(rawBytes)
		fmt.Printf("* ThriftBinary +%d bytes\n", len(rawBytes))
	}

	fmt.Printf("** ThriftBinary %d bytes\n", totalBytesCount)
}
