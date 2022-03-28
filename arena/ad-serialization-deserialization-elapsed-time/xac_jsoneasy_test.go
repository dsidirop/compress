package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"github.com/mailru/easyjson"
)

var jsoneasyIteratorItem interface{}

func Test___SerializationDeserializationElapsedTime___JsonEasy(t *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	newitem := arena.ItemSerdeable(nil)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := datasource[i%datasourceArrayLength]

		bytes, err := easyjson.Marshal(x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		newitem = x.NewEmptyItem()
		err = easyjson.Unmarshal(bytes, newitem)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** JSONEasy %d nanoseconds\n", int64(averageElapsedTime))

	jsoneasyIteratorItem = newitem
}
