package serialization_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationPerformance___GoAvro(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		_, err := avro.Marshal(arena.Schemas.GoAvro, &x)
		if err != nil {
			panic(err)
		}
	}
}
