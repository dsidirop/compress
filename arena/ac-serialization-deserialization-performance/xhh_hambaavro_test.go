package serialization_deserialization_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationDeserializationPerformance___HambaAvro(b *testing.B) {
	datasourceArrayLength := len(arena.MainDatasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		bytes, err := avro.Marshal(x.HambaAvroSchema, x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		newitem := x.NewEmptyItem()
		err = avro.Unmarshal(x.HambaAvroSchema, bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
