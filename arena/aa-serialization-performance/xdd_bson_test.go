package serialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

var bsonRawBytes []byte

func Benchmark___SerializationPerformance___Bson(b *testing.B) {
	err := error(nil)
	rawbytes := []byte(nil)
	datasourceArrayLength := len(arena.MainDatasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		rawbytes, err = bson.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	bsonRawBytes = rawbytes
}
