package decompression_with_deserialization_performance

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___DecompressionAndDeserializationPerformance___Json(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			compressedAndSerializedDatasource := [][]byte{} //first serialize and compress
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				jsonBytes, err := json.Marshal(x)
				if err != nil {
					panic(err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(jsonBytes)
				if err != nil {
					panic(err)
				}

				compressedAndSerializedDatasource = append(compressedAndSerializedDatasource, compressedAndSerializedBytes)
			}

			bench.ResetTimer() //vital
			for i := 0; i < bench.N; i++ {
				x := compressedAndSerializedDatasource[i%datasourceArrayLength] //and now we deserialize and decompress

				serializedBytes, err := test.DecompressionCallback(x)
				if err != nil {
					panic(err)
				}

				fooitem := arena.FooItem{}
				err = json.Unmarshal(serializedBytes, &fooitem)
				if err != nil {
					panic(err)
				}
			}
		})
	}
}
