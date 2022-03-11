package serialization_eventual_message_size_footprint

import (
	"context"
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___ThriftCompact(t *testing.T) {
	ctx := context.TODO()
	item := arena.SpecialDatasourcesForIDLMechanisms.Thrift[0]
	thriftCompactSerializer := arena.NewThriftCompactSerializer()

	rawBytes, err := thriftCompactSerializer.Write(ctx, item.Item)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	fmt.Printf("** ThriftCompact %d bytes\n", len(rawBytes))
}
