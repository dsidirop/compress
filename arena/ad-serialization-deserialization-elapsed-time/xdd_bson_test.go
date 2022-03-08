package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

func Test___SerializationDeserializationElapsedTime___Bson(t *testing.T) {
	datasourceArrayLength := len(arena.Datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		bytes, err := bson.Marshal(x)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		y := &arena.FooItem{}
		err = bson.Unmarshal(bytes, y)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** BSON %d nanoseconds\n", int64(averageElapsedTime))
}
