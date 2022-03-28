package decompression_with_deserialization_performance

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/klauspost/compress/arena"
)

var jsonIteratorNewItem interface{}

func Benchmark___DecompressionAndDeserializationPerformance___JsonIterator(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		newitem := arena.ItemSerdeable(nil)

		b.Run(test.Desc, func(bench *testing.B) {
			compressedAndSerializedDatasource := [][]byte{} //first serialize and compress

			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				jsonBytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(jsonBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedDatasource = append(compressedAndSerializedDatasource, compressedAndSerializedBytes)
			}

			bench.ResetTimer() //vital
			for iteration := 0; iteration < bench.N; iteration++ {
				i := iteration % datasourceArrayLength
				x := compressedAndSerializedDatasource[i] //and now we deserialize and decompress
				mainItemSpec := arena.MainDatasource[i]

				serializedBytes, err := test.DecompressionCallback(x)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				newitem = mainItemSpec.NewEmptyItem()
				err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(serializedBytes, newitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})

		jsonIteratorNewItem = newitem
	}
}
