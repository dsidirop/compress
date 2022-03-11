package serialization_deserialization_performance

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationDeserializationPerformance___Cbor(b *testing.B) { // https://github.com/fxamacker/cbor
	datasourceArrayLength := len(arena.MainDatasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		bytes, err := cbor.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		newitem := x.NewEmptyItem()
		err = cbor.Unmarshal(bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
