package serialization

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

// these tests are invoked multiple times by the benchmarking framework

func Benchmark___Serialization___Json(t *testing.B) {
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		_, err := json.Marshal(x)
		if err != nil {
			panic(err)
		}
	}
}
