package serialization_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

var hambaavroRawBytes []byte

func Benchmark___SerializationPerformance___HambaAvro(b *testing.B) {
	err := error(nil)
	rawbytes := []byte(nil)

	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		rawbytes, err = avro.Marshal(x.HambaAvroSchema, x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	hambaavroRawBytes = rawbytes
}
