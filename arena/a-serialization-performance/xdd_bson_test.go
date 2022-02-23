package serialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

func Benchmark___SerializationPerformance___Bson(t *testing.B) {
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		_, err := bson.Marshal(x)
		if err != nil {
			panic(err)
		}
	}
}
