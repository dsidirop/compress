package serialization_eventual_message_size_footprint

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Test___SerializationMessageSizeFootprint___Msgp(t *testing.T) {
	datasourceArrayLength := len(arena.MainDatasource)

	totalBytesCount := 0
	for i := 0; i < datasourceArrayLength; i++ {
		x := arena.MainDatasource[i]

		buf := &bytes.Buffer{}
		err := msgp.Encode(buf, x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		totalBytesCount += len(buf.Bytes())
	}

	fmt.Printf("** Msgp %d bytes\n", totalBytesCount)
}
