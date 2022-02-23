package serialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

func Benchmark___SerializationPerformance___MessagePack(t *testing.B) {
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		_, err := msgpack.Marshal(x)
		if err != nil {
			panic(err)
		}
	}
}
