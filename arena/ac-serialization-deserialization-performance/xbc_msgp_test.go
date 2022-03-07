package serialization_deserialization_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Benchmark___SerializationDeserializationPerformance___Msgp(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		buf := &bytes.Buffer{}
		err := msgp.Encode(buf, &x)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		fooitem := &arena.FooItem{}
		err = msgp.Decode(buf, fooitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
