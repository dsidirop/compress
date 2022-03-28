package serialization_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/mailru/easyjson"
)

var jsoneasyIteratorNewItem interface{}

func Benchmark___SerializationDeserializationPerformance___JsonEasy(b *testing.B) {
	newitem := arena.ItemSerdeable(nil)
	datasourceArrayLength := len(arena.MainDatasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		bytes, err := easyjson.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		newitem = x.NewEmptyItem()
		err = easyjson.Unmarshal(bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	jsoneasyIteratorNewItem = newitem
}
