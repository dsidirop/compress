package serialization_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

func Test___SerializationMessageSizeFootprint___Bson(t *testing.T) {
	datasourceArrayLength := len(arena.MainDatasource)

	totalBytesCount := 0
	for i := 0; i < datasourceArrayLength; i++ {
		x := arena.MainDatasource[i]

		rawBytes, err := bson.Marshal(x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		totalBytesCount += len(rawBytes)
	}

	fmt.Printf("** BSON %d bytes\n", totalBytesCount)
}
