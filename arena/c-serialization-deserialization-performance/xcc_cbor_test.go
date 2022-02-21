package serialization

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationDeserializationPerformance___Cbor(t *testing.B) { // https://github.com/fxamacker/cbor
	y := arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		bytes, err := cbor.Marshal(x)
		if err != nil {
			panic(err)
		}

		err = cbor.Unmarshal(bytes, &y)
		if err != nil {
			panic(err)
		}
	}
}
