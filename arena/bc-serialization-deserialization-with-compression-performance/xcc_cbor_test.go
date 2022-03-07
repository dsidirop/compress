package serialization_deserialization_performance

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationDeserializationWithCompressionPerformance___Cbor(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer() //vital

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				serializedBytes, err := cbor.Marshal(x)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytes)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				decompressedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				fooitem := &arena.FooItem{}
				err = cbor.Unmarshal(decompressedBytes, fooitem)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
