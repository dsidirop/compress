package serialization_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

var bsonDummyVariable interface{}

func Benchmark___SerializationDeserializationWithCompressionPerformance___Bson(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer() //vital

			newitem := arena.ItemSerdeable(nil) //0 always record the result to prevent the compiler eliminating the function call
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				serializedBytes, err := bson.Marshal(x.Item)
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

				newitem = x.NewEmptyItem()
				err = bson.Unmarshal(decompressedBytes, newitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}

			bsonDummyVariable = newitem
		})
	}
}
