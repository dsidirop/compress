package serialization_performance

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

var jsonRawBytes []byte

func Benchmark___SerializationPerformance___Json(b *testing.B) {
	err := error(nil)
	rawbytes := []byte(nil)
	datasourceArrayLength := len(arena.MainDatasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		rawbytes, err = json.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	jsonRawBytes = rawbytes
}
