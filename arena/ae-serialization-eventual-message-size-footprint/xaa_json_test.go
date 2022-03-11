package serialization_eventual_message_size_footprint

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___Json(t *testing.T) {
	x := arena.MainDatasource[0]

	rawBytes, err := json.Marshal(x.Item)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	fmt.Printf("** JSON %d bytes\n", len(rawBytes))
}
