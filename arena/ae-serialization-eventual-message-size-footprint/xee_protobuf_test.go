package serialization_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Test___SerializationMessageSizeFootprint___Protobuf(t *testing.T) {
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Protobuf
	datasourceArrayLength := len(datasource)

	totalBytesCount := 0
	for i := 0; i < datasourceArrayLength; i++ {
		x := datasource[i]

		rawBytes, err := proto.Marshal(x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		totalBytesCount += len(rawBytes)
	}

	fmt.Printf("** ProtoBuf %d bytes\n", totalBytesCount)
}
