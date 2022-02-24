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
	thriftCompactSerializer := arena.NewThriftCompactSerializer()

	x := datasource[0]

	rawBytes, err := thriftCompactSerializer.Write(ctx, x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("** ThriftCompact %d bytes\n", len(rawBytes))
}
