package serialization

import (
	"log"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___Cbor(t *testing.T) {
	x := arena.Datasource[0]

	rawBytes, err := cbor.Marshal(x)
	if err != nil {
		panic(err)
	}

	log.Printf("** CBOR %d bytes\n", len(rawBytes))
}
