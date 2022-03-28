package serialization_performance

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/klauspost/compress/arena"
)

var jsonIteratorRawBytes []byte

func Benchmark___SerializationPerformance___JsonIterator(b *testing.B) {
	err := error(nil)
	rawbytes := []byte(nil)
	datasourceArrayLength := len(arena.MainDatasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		rawbytes, err = jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	jsonIteratorRawBytes = rawbytes
}
