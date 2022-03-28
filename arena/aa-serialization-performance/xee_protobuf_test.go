package serialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

var protobufRawBytes []byte

func Benchmark___SerializationPerformance___Protobuf(b *testing.B) {
	err := error(nil)
	rawbytes := []byte(nil)

	datasource := arena.SpecialDatasourcesForIDLMechanisms.Protobuf
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		rawbytes, err = proto.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	protobufRawBytes = rawbytes
}
