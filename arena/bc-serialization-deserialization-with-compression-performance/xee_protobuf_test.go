package serialization_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

var protobufDummyVariable interface{}

func Benchmark___SerializationDeserializationWithCompressionPerformance___Protobuf(b *testing.B) {
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Protobuf
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer() //vital

			newitem := arena.ItemSerdeable(nil)
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				serializedBytes, err := proto.Marshal(x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				decompressedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				newitem := x.NewEmptyProtobufItem()
				err = proto.Unmarshal(decompressedBytes, newitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}

			protobufDummyVariable = newitem
		})
	}
}
