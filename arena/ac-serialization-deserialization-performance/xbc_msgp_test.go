package serialization_deserialization_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

var msgpNewItem interface{}

func Benchmark___SerializationDeserializationPerformance___Msgp(b *testing.B) {
	newitem := arena.ItemSerdeable(nil)
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		buf := &bytes.Buffer{}
		err := msgp.Encode(buf, x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		newitem = x.NewEmptyItem()
		err = msgp.Decode(buf, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	msgpNewItem = newitem
}
