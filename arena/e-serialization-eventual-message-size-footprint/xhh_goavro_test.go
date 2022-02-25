package serialization_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___GoAvro(t *testing.T) {
	datasource := arena.Datasource

	x := datasource[0]

	rawBytes, err := avro.Marshal(arena.Schemas.GoAvro, &x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("** GoAvro %d bytes\n", len(rawBytes))
}
