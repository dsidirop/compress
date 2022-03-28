package deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Benchmark___Deserialization___Protobuf(b *testing.B) {
	newitem := proto.Message(nil)
	datasource := arena.SerializedDataSources.Protobuf
	datasourceArrayLength := len(datasource)

	b.ResetTimer()
	for iterator := 0; iterator < b.N; iterator++ {
		i := iterator % datasourceArrayLength

		newitem = arena.MainDatasource[i].NewEmptyProtobufItem()
		rawbytes := datasource[i]

		err := proto.Unmarshal(rawbytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}
}
