package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationDeserializationElapsedTime___Cbor(t *testing.T) { // https://github.com/fxamacker/cbor
	y := arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		bytes, err := cbor.Marshal(x)
		if err != nil {
			panic(err)
		}

		err = cbor.Unmarshal(bytes, &y)
		if err != nil {
			panic(err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** CBOR %d nanoseconds\n", int64(averageElapsedTime))
}
