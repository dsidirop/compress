package decompression_with_deserialization_performance

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Benchmark___DecompressionAndDeserializationPerformance___Cbor(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			compressedAndSerializedDatasource := [][]byte{} //first serialize and compress
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				serializedBytes, err := cbor.Marshal(x.Item)
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

				newitem := mainItemSpec.NewEmptyItem()
				err = cbor.Unmarshal(decompressedSerializedBytes, newitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
