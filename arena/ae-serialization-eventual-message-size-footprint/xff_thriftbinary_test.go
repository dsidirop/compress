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
	x := arena.SpecialDatasourcesForIDLMechanisms.Thrift[0]
	thriftBinarySerializer := thrift.NewTSerializer() //binary serializer

	rawBytes, err := thriftBinarySerializer.Write(ctx, x.Item)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	fmt.Printf("** ThriftBinary %d bytes\n", len(rawBytes))
}
