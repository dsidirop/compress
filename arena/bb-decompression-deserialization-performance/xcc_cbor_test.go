package decompression_with_deserialization_performance

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Benchmark___DecompressionAndDeserializationPerformance___Cbor(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		// if test.Desc != "Deflate" {
		// 	continue
		// }

		b.Run(test.Desc, func(bench *testing.B) {
			compressedAndSerializedDatasource := [][]byte{} //first serialize and compress
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				serializedBytes, err := cbor.Marshal(x)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytes)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				compressedAndSerializedDatasource = append(compressedAndSerializedDatasource, compressedAndSerializedBytes)
			}

			bench.ResetTimer() //vital
			for i := 0; i < bench.N; i++ {
				x := compressedAndSerializedDatasource[i%datasourceArrayLength] //and now we deserialize and decompress

				decompressedSerializedBytes, err := test.DecompressionCallback(x)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				item := &arena.FooItem{}
				err = cbor.Unmarshal(decompressedSerializedBytes, item)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
