package serialization_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

func Test___SerializationMessageSizeFootprint___MessagePack(t *testing.T) {
	x := arena.Datasource[0]

	rawBytes, err := msgpack.Marshal(x)
	if err != nil {
		b.Fatalf("Error: %s", err)
	}

	fmt.Printf("** MessagePack %d bytes\n", len(rawBytes))
}
