package deserialization_performance

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___Deserialization___Json(b *testing.B) {
	datasource := arena.SerializedDataSources.Json
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		i := iteration % datasourceArrayLength

		bytes := datasource[i]
		newitem := arena.MainDatasource[i].NewEmptyItem()

		err := json.Unmarshal(bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
