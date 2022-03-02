package serialization_deserialization_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Benchmark___SerializationDeserializationPerformance___Msgp(b *testing.B) {
	buf := &bytes.Buffer{}
	fooitem := &arena.FooItem{}
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		err := msgp.Encode(buf, &x)
		if err != nil {
			panic(err)
		}

		err = msgp.Decode(buf, fooitem)
		if err != nil {
			panic(err)
		}
	}
}
