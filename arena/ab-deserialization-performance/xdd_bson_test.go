package deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

func Benchmark___Deserialization___Bson(b *testing.B) {
	newitem := arena.ItemSerdeable(nil)
	datasource := arena.SerializedDataSources.Bson
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for iterator := 0; iterator < b.N; iterator++ {
		i := iterator % datasourceArrayLength

		newitem = arena.MainDatasource[i].NewEmptyItem()
		rawbytes := datasource[i]

		err := bson.Unmarshal(rawbytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
