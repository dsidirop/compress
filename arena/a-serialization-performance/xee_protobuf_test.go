package serialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Benchmark___SerializationPerformance___Protobuf(t *testing.B) {
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		x := arena.DatasourceForProtobuf[i%datasourceArrayLength]

		_, err := proto.Marshal(x)
		if err != nil {
			panic(err)
		}
	}
}
