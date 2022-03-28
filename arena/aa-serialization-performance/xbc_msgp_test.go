package serialization_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

var msgpRawBytes []byte

func Benchmark___SerializationPerformance___Msgp(b *testing.B) {
	err := error(nil)
	bytesbuffer := bytes.Buffer{}

	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		bytesbuffer = bytes.Buffer{} //keep here
		err = msgp.Encode(&bytesbuffer, x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	msgpRawBytes = bytesbuffer.Bytes()
}
