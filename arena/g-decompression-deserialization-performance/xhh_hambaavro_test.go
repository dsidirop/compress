package decompression_with_deserialization_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Benchmark___DecompressionAndDeserializationPerformance___HambaAvro(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			compressedAndSerializedDatasource := [][]byte{} //first serialize and compress
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				serializedBytes, err := avro.Marshal(arena.Schemas.GoHambaAvro, &x)
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
			for i := 0; i < bench.N; i++ {
				x := compressedAndSerializedDatasource[i%datasourceArrayLength] //and now we deserialize and decompress

				decompressedSerializedBytes, err := test.DecompressionCallback(x)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				y := &arena.FooItem{}
				err = avro.Unmarshal(arena.Schemas.GoHambaAvro, decompressedSerializedBytes, y)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
