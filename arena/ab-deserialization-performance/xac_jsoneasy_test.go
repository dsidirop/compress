package deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/mailru/easyjson"
)

func Benchmark___Deserialization___JsonEasy(b *testing.B) {
	newitem := arena.ItemSerdeable(nil)
	datasource := arena.SerializedDataSources.Json
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		i := iteration % datasourceArrayLength

		bytes := datasource[i]
		newitem = arena.MainDatasource[i].NewEmptyItem()

		err := easyjson.Unmarshal(bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
