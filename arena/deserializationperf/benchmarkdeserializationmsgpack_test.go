package arena

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

func Benchmark___Deserialization___MessagePack(t *testing.B) {
	item := arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		bytes := arena.DatasourceSerializedMessagePack[i%datasourceArrayLength]

		err := msgpack.Unmarshal(bytes, &item)
		if err != nil {
			panic(err)
		}
	}
}
