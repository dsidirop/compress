package decompression_with_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

func Benchmark___DecompressionAndDeserializationPerformance___MessagePack(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			compressedAndSerializedDatasource := [][]byte{} //first serialize and compress
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				msgpackBytes, err := msgpack.Marshal(x)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(msgpackBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedDatasource = append(compressedAndSerializedDatasource, compressedAndSerializedBytes)
			}

			bench.ResetTimer() //vital
			for i := 0; i < bench.N; i++ {
				x := compressedAndSerializedDatasource[i%datasourceArrayLength] //and now we deserialize and decompress

				serializedBytes, err := test.DecompressionCallback(x)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				fooitem := arena.FooItem{}
				err = msgpack.Unmarshal(serializedBytes, &fooitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
