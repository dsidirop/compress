package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationDeserializationElapsedTime___GoAvro(t *testing.T) {
	y := &arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		goAvroBytes, err := avro.Marshal(arena.Schemas.GoAvro, &x)
		if err != nil {
			panic(err)
		}

		err = avro.Unmarshal(arena.Schemas.GoAvro, goAvroBytes, y)
		if err != nil {
			panic(err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** GoAvro %d nanoseconds\n", int64(averageElapsedTime))
}
