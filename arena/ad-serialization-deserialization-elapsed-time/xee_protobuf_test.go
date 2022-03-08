package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Test___SerializationDeserializationElapsedTime___Protobuf(t *testing.T) {
	datasourceArrayLength := len(arena.Datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := arena.SpecialDatasourcesForIDLMechanisms.Protobuf[i%datasourceArrayLength]

		bytes, err := proto.Marshal(x)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		y := &arena.PBFooItem{}
		err = proto.Unmarshal(bytes, y)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** ProtoBuf %d nanoseconds\n", int64(averageElapsedTime))
}
