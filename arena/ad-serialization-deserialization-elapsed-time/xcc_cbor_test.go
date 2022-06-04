package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationDeserializationElapsedTime___Cbor(t *testing.T) { // https://github.com/fxamacker/cbor
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := datasource[i%datasourceArrayLength]

		bytes, err := cbor.Marshal(x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		newitem := x.NewEmptyItem()
		err = cbor.Unmarshal(bytes, newitem)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** CBOR %d nanoseconds\n", int64(averageElapsedTime))
}
