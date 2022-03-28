package deserialization_performance

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Benchmark___Deserialization___Cbor(b *testing.B) {
	newitem := arena.ItemSerdeable(nil)
	datasource := arena.SerializedDataSources.Cbor
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		i := iteration % datasourceArrayLength

		newitem = arena.MainDatasource[i].NewEmptyItem()
		rawbytes := datasource[i]

		err := cbor.Unmarshal(rawbytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
