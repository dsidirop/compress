package serialization

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationPerformance___Json(t *testing.B) {
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		_, err := json.Marshal(x)
		if err != nil {
			panic(err)
		}
	}
}
