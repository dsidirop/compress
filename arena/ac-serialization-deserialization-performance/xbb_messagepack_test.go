package serialization_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

func Benchmark___SerializationDeserializationPerformance___MessagePack(b *testing.B) {
	y := arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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
