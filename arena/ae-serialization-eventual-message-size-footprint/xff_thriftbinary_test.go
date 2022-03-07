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
	thriftBinarySerializer := thrift.NewTSerializer() //binary serializer

	x := datasource[0]

	rawBytes, err := thriftBinarySerializer.Write(ctx, x)
	if err != nil {
		b.Fatalf("Error: %s", err)
	}

	fmt.Printf("** ThriftBinary %d bytes\n", len(rawBytes))
}
