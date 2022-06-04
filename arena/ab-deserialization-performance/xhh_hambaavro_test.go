package deserialization_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Benchmark___DeserializationPerformance___HambaAvro(b *testing.B) {
	datasource := arena.SerializedDataSources.GoHambaAvro
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for iterator := 0; iterator < b.N; iterator++ {
		i := iterator % datasourceArrayLength

		mainDatasourceEntry := arena.MainDatasource[i]

		schema := mainDatasourceEntry.HambaAvroSchema
		newitem := mainDatasourceEntry.NewEmptyItem()
		rawbytes := datasource[i]

		err := avro.Unmarshal(schema, rawbytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
