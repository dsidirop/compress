package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationDeserializationElapsedTime___HambaAvro(t *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	newitem := arena.ItemSerdeable(nil)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := datasource[i%datasourceArrayLength]

		goAvroBytes, err := avro.Marshal(x.HambaAvroSchema, x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		newitem = x.NewEmptyItem()
		err = avro.Unmarshal(x.HambaAvroSchema, goAvroBytes, newitem)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** HambaAvro %d nanoseconds\n", int64(averageElapsedTime))
}
