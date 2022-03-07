package serialization_deserialization_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationDeserializationPerformance___HambaAvro(b *testing.B) {
	y := arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.Datasource[i%datasourceArrayLength]

		bytes, err := avro.Marshal(arena.Schemas.GoHambaAvro, x)
		if err != nil {
			panic(err)
		}

		err = avro.Unmarshal(arena.Schemas.GoHambaAvro, bytes, &y)
		if err != nil {
			panic(err)
		}
	}
}
