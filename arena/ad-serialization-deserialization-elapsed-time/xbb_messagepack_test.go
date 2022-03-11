package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

func Test___SerializationDeserializationElapsedTime___MessagePack(t *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := datasource[i%datasourceArrayLength]

		bytes, err := msgpack.Marshal(x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		newitem := x.NewEmptyItem()
		err = msgpack.Unmarshal(bytes, newitem)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** MessagePack %d nanoseconds\n", int64(averageElapsedTime))
}
