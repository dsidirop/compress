package serialization_deserialization_performance

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/klauspost/compress/arena"
)

var jsonIteratorNewItem interface{}

func Benchmark___SerializationDeserializationPerformance___JsonIterator(b *testing.B) {
	newitem := arena.ItemSerdeable(nil)
	datasourceArrayLength := len(arena.MainDatasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		bytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		newitem = x.NewEmptyItem()
		err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	jsonIteratorNewItem = newitem
}
