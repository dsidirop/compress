package serialization_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Benchmark___SerializationDeserializationPerformance___Protobuf(t *testing.B) {
	y := arena.PBFooItem{}
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		x := arena.DatasourceForProtobuf[i%datasourceArrayLength]

		bytes, err := proto.Marshal(x)
		if err != nil {
			panic(err)
		}

		err = proto.Unmarshal(bytes, &y)
		if err != nil {
			panic(err)
		}
	}
}
