package serialization_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

func Test___SerializationMessageSizeFootprint___Bson(t *testing.T) {
	x := arena.MainDatasource[0]

	rawBytes, err := bson.Marshal(x.Item)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	fmt.Printf("** BSON %d bytes\n", len(rawBytes))
}
