package serialization_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___HambaAvro(t *testing.T) {
	x := arena.MainDatasource[0]

	rawBytes, err := avro.Marshal(x.HambaAvroSchema, x.Item)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	fmt.Printf("** HambaAvro %d bytes\n", len(rawBytes))
}
