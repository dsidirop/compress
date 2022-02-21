package arena

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Benchmark___Deserialization___Cbor(t *testing.B) {
	item := arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		bytes := arena.SerializedDataSources.Cbor[i%datasourceArrayLength]

		err := cbor.Unmarshal(bytes, &item)
		if err != nil {
			panic(err)
		}
	}
}
