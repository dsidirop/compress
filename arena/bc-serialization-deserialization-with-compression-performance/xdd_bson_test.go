package serialization_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

func Benchmark___SerializationDeserializationPerformance___Bson(b *testing.B) {
	y := arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		bytes, err := bson.Marshal(x)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		err = bson.Unmarshal(bytes, &y)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
