package serialization

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

// these tests are invoked multiple times by the benchmarking framework

func Benchmark___Serialization___MessagePack(t *testing.B) {
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		_, err := msgpack.Marshal(x)
		if err != nil {
			panic(err)
		}
	}
}
