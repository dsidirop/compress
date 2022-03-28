package serialization_deserialization_performance

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

var jsonNewItem interface{}

func Benchmark___SerializationDeserializationPerformance___Json(b *testing.B) {
	datasourceArrayLength := len(arena.MainDatasource)

	newitem := arena.ItemSerdeable(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		bytes, err := json.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		newitem = x.NewEmptyItem()
		err = json.Unmarshal(bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	jsonNewItem = newitem
}
