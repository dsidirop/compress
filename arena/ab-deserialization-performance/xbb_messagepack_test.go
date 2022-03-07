package deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

func Benchmark___Deserialization___MessagePack(b *testing.B) {
	item := arena.FooItem{}

	datasource := arena.SerializedDataSources.MessagePack
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes := datasource[i%datasourceArrayLength]

		err := msgpack.Unmarshal(bytes, &item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
