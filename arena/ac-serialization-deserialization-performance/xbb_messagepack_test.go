package serialization_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

var messagepackNewItem interface{}

func Benchmark___SerializationDeserializationPerformance___MessagePack(b *testing.B) {
	datasourceArrayLength := len(arena.MainDatasource)

	newitem := arena.ItemSerdeable(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		bytes, err := msgpack.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		newitem = x.NewEmptyItem()
		err = msgpack.Unmarshal(bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	messagepackNewItem = newitem
}
