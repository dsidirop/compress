package serialization_deserialization_elapsed_time

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Test___SerializationDeserializationElapsedTime___Msgp(t *testing.T) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := datasource[i%datasourceArrayLength]

		encodedBytesBuffer := &bytes.Buffer{}
		err := msgp.Encode(encodedBytesBuffer, x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		newitem := x.NewEmptyItem()
		err = msgp.Decode(encodedBytesBuffer, newitem)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** Msgp %d nanoseconds\n", int64(averageElapsedTime))
}
