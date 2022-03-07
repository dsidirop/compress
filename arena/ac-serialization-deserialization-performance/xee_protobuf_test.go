package serialization_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Benchmark___SerializationDeserializationPerformance___Protobuf(b *testing.B) {
	y := arena.PBFooItem{}
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Protobuf
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := datasource[i%datasourceArrayLength]

		bytes, err := proto.Marshal(x)
		if err != nil {
			panic(err)
		}

		err = proto.Unmarshal(bytes, &y)
		if err != nil {
			panic(err)
		}
	}
}
