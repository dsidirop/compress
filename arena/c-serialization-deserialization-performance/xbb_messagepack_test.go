package serialization

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

func Benchmark___SerializationDeserializationPerformance___MessagePack(t *testing.B) {
	y := arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		bytes, err := msgpack.Marshal(x)
		if err != nil {
			panic(err)
		}

		err = msgpack.Unmarshal(bytes, &y)
		if err != nil {
			panic(err)
		}
	}
}
