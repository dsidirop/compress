package serialization_eventual_message_size_footprint

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Test___SerializationMessageSizeFootprint___Msgp(t *testing.T) {
	x := arena.Datasource[0]
	buf := &bytes.Buffer{}

	err := msgp.Encode(buf, &x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("** Msgp %d bytes\n", len(buf.Bytes()))
}
