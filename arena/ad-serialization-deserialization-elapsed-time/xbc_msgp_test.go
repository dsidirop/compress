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
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := datasource[i%datasourceArrayLength]

		buf := &bytes.Buffer{}
		err := msgp.Encode(buf, &x)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		fooitem := &arena.FooItem{}
		err = msgp.Decode(buf, fooitem)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** Msgp %d nanoseconds\n", int64(averageElapsedTime))
}
