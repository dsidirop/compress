package decompression_with_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"google.golang.org/protobuf/proto"
)

func Benchmark___DecompressionAndDeserializationPerformance___Protobuf(b *testing.B) {
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Protobuf
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				serializedBytes, err := proto.Marshal(x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				x.Bytes = compressedAndSerializedBytes
			}

			bench.ResetTimer() //vital
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength] //and now we deserialize and decompress

				decompressedSerializedBytes, err := test.DecompressionCallback(x.Bytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				item := x.NewEmptyProtobufItem()
				err = proto.Unmarshal(decompressedSerializedBytes, item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
