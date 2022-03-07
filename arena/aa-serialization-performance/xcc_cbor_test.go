package serialization_performance

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationPerformance___Cbor(b *testing.B) { // https://github.com/fxamacker/cbor
	datasourceArrayLength := len(arena.Datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		_, err := cbor.Marshal(x)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
