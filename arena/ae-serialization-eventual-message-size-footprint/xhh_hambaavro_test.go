package serialization_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___HambaAvro(t *testing.T) {
	datasource := arena.Datasource

	x := datasource[0]

	rawBytes, err := avro.Marshal(arena.Schemas.GoHambaAvro, &x)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	fmt.Printf("** HambaAvro %d bytes\n", len(rawBytes))
}
