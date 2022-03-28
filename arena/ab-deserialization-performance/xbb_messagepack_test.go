package deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

func Benchmark___Deserialization___MessagePack(b *testing.B) {
	newitem := arena.ItemSerdeable(nil)
	datasource := arena.SerializedDataSources.MessagePack
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		i := iteration % datasourceArrayLength

		bytes := datasource[i]
		newitem = arena.MainDatasource[i].NewEmptyItem()

		err := msgpack.Unmarshal(bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
