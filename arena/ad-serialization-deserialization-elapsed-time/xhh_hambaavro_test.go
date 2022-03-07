package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationDeserializationElapsedTime___HambaAvro(t *testing.T) {
	y := &arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		goAvroBytes, err := avro.Marshal(arena.Schemas.GoHambaAvro, &x)
		if err != nil {
			panic(err)
		}

		err = avro.Unmarshal(arena.Schemas.GoHambaAvro, goAvroBytes, y)
		if err != nil {
			panic(err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** HambaAvro %d nanoseconds\n", int64(averageElapsedTime))
}
