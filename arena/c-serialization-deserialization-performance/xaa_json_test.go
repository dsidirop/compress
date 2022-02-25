package serialization_deserialization_performance

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationDeserializationPerformance___Json(b *testing.B) {
	y := arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		bytes, err := json.Marshal(x)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &y)
		if err != nil {
			panic(err)
		}
	}
}
