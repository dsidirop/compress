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
	buf := &bytes.Buffer{}
	fooitem := &arena.FooItem{}
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := datasource[i%datasourceArrayLength]

		err := msgp.Encode(buf, &x)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		err = msgp.Decode(buf, fooitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** Msgp %d nanoseconds\n", int64(averageElapsedTime))
}
