package serialization_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Benchmark___SerializationDeserializationPerformance___Protobuf(b *testing.B) {
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Protobuf
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		bytes, err := proto.Marshal(x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		newitem := x.NewEmptyProtobufItem()
		err = proto.Unmarshal(bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
