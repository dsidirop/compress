package deserialization_performance

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/klauspost/compress/arena"
)

func Benchmark___Deserialization___JsonIterator(b *testing.B) {
	newitem := arena.ItemSerdeable(nil)
	datasource := arena.SerializedDataSources.Json
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		i := iteration % datasourceArrayLength

		bytes := datasource[i]
		newitem = arena.MainDatasource[i].NewEmptyItem()

		err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
