package deserialization_performance

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___Deserialization___Json(b *testing.B) {
	item := arena.FooItem{}

	datasource := arena.SerializedDataSources.Json
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes := datasource[i%datasourceArrayLength]

		err := json.Unmarshal(bytes, &item)
		if err != nil {
			panic(err)
		}
	}
}
