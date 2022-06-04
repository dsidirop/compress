package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Test___SerializationDeserializationElapsedTime___Protobuf(t *testing.T) {
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Protobuf
	datasourceArrayLength := len(datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := datasource[i%datasourceArrayLength]

		bytes, err := proto.Marshal(x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		newitem := x.NewEmptyProtobufItem()
		err = proto.Unmarshal(bytes, newitem)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** ProtoBuf %d nanoseconds\n", int64(averageElapsedTime))
}
