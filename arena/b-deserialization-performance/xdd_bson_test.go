package deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

func Benchmark___Deserialization___Bson(t *testing.B) {
	item := arena.FooItem{}

	datasource := arena.SerializedDataSources.Bson
	datasourceArrayLength := len(datasource)

	for i := 0; i < t.N; i++ {
		bytes := datasource[i%datasourceArrayLength]

		err := bson.Unmarshal(bytes, &item)
		if err != nil {
			panic(err)
		}
	}
}
