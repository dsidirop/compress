package deserialization_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Benchmark___Deserialization___Msgp(b *testing.B) {
	newitem := arena.ItemSerdeable(nil)
	datasource := arena.SerializedDataSources.Msgp
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		i := iteration % datasourceArrayLength

		newitem = arena.MainDatasource[i].NewEmptyItem()
		byteBuffer := bytes.NewBuffer(datasource[i])

		err := msgp.Decode(byteBuffer, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
