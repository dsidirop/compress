package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

func Test___SerializationDeserializationElapsedTime___MessagePack(t *testing.T) {
	datasourceArrayLength := len(arena.Datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		bytes, err := msgpack.Marshal(x)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		y := &arena.FooItem{}
		err = msgpack.Unmarshal(bytes, y)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** MessagePack %d nanoseconds\n", int64(averageElapsedTime))
}
