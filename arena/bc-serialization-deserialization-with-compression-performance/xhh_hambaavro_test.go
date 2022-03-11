package serialization_deserialization_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationDeserializationWithCompressionPerformance___HambaAvro(rootBench *testing.B) {
	datasourceArrayLength := len(arena.MainDatasource)

	for _, test := range arena.AllCompressionTestCases {
		rootBench.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer() //vital

			for i := 0; i < bench.N; i++ {
				x := arena.MainDatasource[i%datasourceArrayLength]

				bytes, err := avro.Marshal(x.HambaAvroSchema, x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(bytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				decompressedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				newitem := x.NewEmptyItem()
				err = avro.Unmarshal(x.HambaAvroSchema, decompressedBytes, newitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
