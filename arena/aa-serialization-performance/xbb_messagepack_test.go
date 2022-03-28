package serialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

var messagepackRawBytes []byte

func Benchmark___SerializationPerformance___MessagePack(b *testing.B) {
	err := error(nil)
	rawbytes := []byte(nil)
	datasourceArrayLength := len(arena.MainDatasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		rawbytes, err = msgpack.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	messagepackRawBytes = rawbytes
}
