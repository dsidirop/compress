package deserialization_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Benchmark___Deserialization___Msgp(b *testing.B) {
	fooitem := &arena.FooItem{}
	datasource := arena.SerializedDataSources.Msgp
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rawBytes := datasource[i%datasourceArrayLength]

		byteBuffer := bytes.NewBuffer(rawBytes) // unfortunate necessity

		err := msgp.Decode(byteBuffer, fooitem)
		if err != nil {
			panic(err)
		}
	}
}
