package deserialization_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Benchmark___DeserializationPerformance___GoAvro(b *testing.B) {
	y := &arena.FooItem{}
	datasource := arena.SerializedDataSources.GoAvro
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		err := avro.Unmarshal(arena.Schemas.GoAvro, x, y)
		if err != nil {
			panic(err)
		}
	}
}
