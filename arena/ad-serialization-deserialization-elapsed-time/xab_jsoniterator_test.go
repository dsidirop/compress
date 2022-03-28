package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/klauspost/compress/arena"
)

var jsonIteratorItem interface{}

func Test___SerializationDeserializationElapsedTime___JsonIterator(t *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	newitem := arena.ItemSerdeable(nil)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := datasource[i%datasourceArrayLength]

		bytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		newitem = x.NewEmptyItem()
		err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(bytes, newitem)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** JSONIterator %d nanoseconds\n", int64(averageElapsedTime))

	jsonIteratorItem = newitem
}
