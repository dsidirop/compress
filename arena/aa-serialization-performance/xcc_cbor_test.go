package serialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
)

var cborRawBytes []byte

func Benchmark___SerializationPerformance___Cbor(b *testing.B) { // https://github.com/fxamacker/cbor
	err := error(nil)
	rawbytes := []byte(nil)

	datasourceArrayLength := len(arena.MainDatasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		rawbytes, err = arena.CborStandardEncoder.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	cborRawBytes = rawbytes
}
