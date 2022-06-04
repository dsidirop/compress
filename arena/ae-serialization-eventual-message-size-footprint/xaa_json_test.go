package serialization_eventual_message_size_footprint

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___Json(t *testing.T) {
	datasourceArrayLength := len(arena.MainDatasource)

	totalBytesCount := 0
	for i := 0; i < datasourceArrayLength; i++ {
		x := arena.MainDatasource[i]

		rawBytes, err := json.Marshal(x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		totalBytesCount += len(rawBytes)
	}

	fmt.Printf("** JSON %d bytes\n", totalBytesCount)
}
