package deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

func Benchmark___Deserialization___Bson(b *testing.B) {
	item := arena.FooItem{}

	datasource := arena.SerializedDataSources.Bson
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes := datasource[i%datasourceArrayLength]

		err := bson.Unmarshal(bytes, &item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
