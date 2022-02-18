package serialization

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___Json(t *testing.T) {
	x := arena.Datasource[0]

	rawBytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	log.Printf("** JSON %d bytes\n", len(rawBytes))
}
