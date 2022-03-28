package serialization_eventual_message_size_footprint

import (
	"fmt"
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___HambaAvro(t *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	totalBytesCount := 0
	for i := 0; i < datasourceArrayLength; i++ {
		x := datasource[i]

		rawBytes, err := avro.Marshal(x.HambaAvroSchema, x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		totalBytesCount += len(rawBytes)
		fmt.Printf("* HambaAvro +%d bytes\n", len(rawBytes))
	}

	fmt.Printf("** HambaAvro %d bytes\n", totalBytesCount)
}
