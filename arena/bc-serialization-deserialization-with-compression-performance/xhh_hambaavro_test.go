package serialization_deserialization_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationDeserializationWithCompressionPerformance___HambaAvro(rootBench *testing.B) {
	datasourceArrayLength := len(arena.Datasource)

	for _, test := range arena.AllCompressionTestCases {
		rootBench.Run(test.Desc, func(subbench *testing.B) {
			subbench.ResetTimer() //vital

			for i := 0; i < subbench.N; i++ {
				x := arena.Datasource[i%datasourceArrayLength]

				bytes, err := avro.Marshal(arena.Schemas.GoHambaAvro, x)
				if err != nil {
					subbench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(bytes)
				if err != nil {
					subbench.Fatalf("Error: %s", err)
				}

				decompressedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					subbench.Fatalf("Error: %s", err)
				}

				fooitem := &arena.FooItem{}
				err = avro.Unmarshal(arena.Schemas.GoHambaAvro, decompressedBytes, fooitem)
				if err != nil {
					subbench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
