package serialization_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Test___SerializationMessageSizeFootprint___Protobuf(t *testing.T) {
	x := arena.ConvertFooItemToPBFooItem(arena.Datasource[0])

	rawBytes, err := proto.Marshal(&x)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	fmt.Printf("** ProtoBuf %d bytes\n", len(rawBytes))
}
