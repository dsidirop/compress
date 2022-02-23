package serialization_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

func Benchmark___SerializationDeserializationPerformance___Bson(t *testing.B) {
	y := arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		bytes, err := bson.Marshal(x)
		if err != nil {
			panic(err)
		}

		err = bson.Unmarshal(bytes, &y)
		if err != nil {
			panic(err)
		}
	}
}
