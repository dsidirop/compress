package decompression_with_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

var bsonNewItem interface{}

func Benchmark___DecompressionAndDeserializationPerformance___Bson(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		newitem := arena.ItemSerdeable(nil)

		b.Run(test.Desc, func(bench *testing.B) {
			compressedAndSerializedDatasource := [][]byte{} //first serialize and compress
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				serializedBytes, err := bson.Marshal(x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedDatasource = append(compressedAndSerializedDatasource, compressedAndSerializedBytes)
			}

			bench.ResetTimer() //vital
			for iterator := 0; iterator < bench.N; iterator++ {
				i := iterator % datasourceArrayLength
				x := compressedAndSerializedDatasource[i] //and now we deserialize and decompress
				mainItemSpec := arena.MainDatasource[i]

				decompressedSerializedBytes, err := test.DecompressionCallback(x)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				newitem = mainItemSpec.NewEmptyItem()
				err = bson.Unmarshal(decompressedSerializedBytes, newitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})

		bsonNewItem = newitem
	}
}
