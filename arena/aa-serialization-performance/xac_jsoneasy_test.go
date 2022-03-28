package serialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/mailru/easyjson"
)

var jsonEasyRawBytes []byte

func Benchmark___SerializationPerformance___JsonEasy(b *testing.B) {
	err := error(nil)
	rawbytes := []byte(nil)
	datasourceArrayLength := len(arena.MainDatasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		rawbytes, err = easyjson.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	jsonRawBytes = rawbytes
}
