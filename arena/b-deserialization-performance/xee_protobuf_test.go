package deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Benchmark___Deserialization___Protobuf(b *testing.B) {
	item := arena.PBFooItem{}
	datasource := arena.SerializedDataSources.Protobuf
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes := datasource[i%datasourceArrayLength]

		err := proto.Unmarshal(bytes, &item)
		if err != nil {
			panic(err)
		}
	}
}
