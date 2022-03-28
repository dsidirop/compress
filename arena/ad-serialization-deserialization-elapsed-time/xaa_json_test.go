package serialization_deserialization_elapsed_time

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
)

var jsonItem interface{}

func Test___SerializationDeserializationElapsedTime___Json(t *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	newitem := arena.ItemSerdeable(nil)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := datasource[i%datasourceArrayLength]

		bytes, err := json.Marshal(x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		newitem = x.NewEmptyItem()
		err = json.Unmarshal(bytes, newitem)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** JSON %d nanoseconds\n", int64(averageElapsedTime))

	jsonItem = newitem
}
