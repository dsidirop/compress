package serialization_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___JsonIterator(t *testing.T) {
	datasourceArrayLength := len(arena.MainDatasource)

	totalBytesCount := 0
	for i := 0; i < datasourceArrayLength; i++ {
		x := arena.MainDatasource[i]

		rawBytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		totalBytesCount += len(rawBytes)
		fmt.Printf("* JSONIterator +%d bytes\n", len(rawBytes))
	}

	fmt.Printf("** JSONIterator %d bytes\n", totalBytesCount)
}
