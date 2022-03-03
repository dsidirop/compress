package serialization_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Benchmark___SerializationPerformance___Msgp(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	b.ResetTimer() //something goes wrong when doing this in this particular test    weird
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		buf := bytes.Buffer{} //keep here
		err := msgp.Encode(&buf, &x)
		if err != nil {
			panic(err)
		}
	}
}
