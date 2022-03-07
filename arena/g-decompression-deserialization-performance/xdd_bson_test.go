package decompression_with_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"go.mongodb.org/mongo-driver/bson"
)

func Benchmark___DecompressionAndDeserializationPerformance___Bson(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			compressedAndSerializedDatasource := [][]byte{} //first serialize and compress
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				serializedBytes, err := bson.Marshal(x)
				if err != nil {
					panic(err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytes)
				if err != nil {
					panic(err)
				}

				compressedAndSerializedDatasource = append(compressedAndSerializedDatasource, compressedAndSerializedBytes)
			}

			bench.ResetTimer() //vital
			for i := 0; i < bench.N; i++ {
				x := compressedAndSerializedDatasource[i%datasourceArrayLength] //and now we deserialize and decompress

				decompressedSerializedBytes, err := test.DecompressionCallback(x)
				if err != nil {
					panic(err)
				}

				item := &arena.FooItem{}
				err = bson.Unmarshal(decompressedSerializedBytes, item)
				if err != nil {
					panic(err)
				}
			}
		})
	}
}
