package deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Benchmark___Deserialization___Protobuf(t *testing.B) {
	item := arena.PBFooItem{}

	datasource := arena.SerializedDataSources.Protobuf
	datasourceArrayLength := len(datasource)

	for i := 0; i < t.N; i++ {
		bytes := datasource[i%datasourceArrayLength]

		err := proto.Unmarshal(bytes, &item)
		if err != nil {
			panic(err)
		}
	}
}
