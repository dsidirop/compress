package decompression_with_deserialization_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

var hambaavroNewItem interface{}

func Benchmark___DecompressionAndDeserializationPerformance___HambaAvro(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		newitem := arena.ItemSerdeable(nil)

		b.Run(test.Desc, func(bench *testing.B) {
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				serializedBytes, err := avro.Marshal(x.HambaAvroSchema, x.Item)
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

				newitem = x.NewEmptyItem()
				err = avro.Unmarshal(x.HambaAvroSchema, decompressedSerializedBytes, newitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})

		hambaavroNewItem = newitem
	}
}
