package deserialization_performance

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Benchmark___Deserialization___Cbor(b *testing.B) {
	item := arena.FooItem{}

	datasource := arena.SerializedDataSources.Cbor
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes := datasource[i%datasourceArrayLength]

		err := cbor.Unmarshal(bytes, &item)
		if err != nil {
			panic(err)
		}
	}
}
