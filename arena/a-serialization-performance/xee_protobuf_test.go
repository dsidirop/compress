package serialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Benchmark___SerializationPerformance___Protobuf(t *testing.B) {
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Protobuf
	datasourceArrayLength := len(datasource)

	for i := 0; i < t.N; i++ {
		x := datasource[i%datasourceArrayLength]

		_, err := proto.Marshal(x)
		if err != nil {
			panic(err)
		}
	}
}
