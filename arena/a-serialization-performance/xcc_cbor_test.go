package serialization

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationPerformance___Cbor(t *testing.B) { // https://github.com/fxamacker/cbor
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		_, err := cbor.Marshal(x)
		if err != nil {
			panic(err)
		}
	}
}
