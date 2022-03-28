package serialization_eventual_message_size_footprint

import (
	"context"
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___ThriftCompact(t *testing.T) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)
	thriftCompactSerializer := arena.NewThriftCompactSerializer()

	totalBytesCount := 0
	for i := 0; i < datasourceArrayLength; i++ {
		x := datasource[i]

		rawBytes, err := thriftCompactSerializer.Write(ctx, x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		totalBytesCount += len(rawBytes)
		fmt.Printf("* ThriftCompact +%d bytes\n", len(rawBytes))
	}

	fmt.Printf("** ThriftCompact %d bytes\n", totalBytesCount)
}
