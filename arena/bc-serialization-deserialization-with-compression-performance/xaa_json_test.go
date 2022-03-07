package serialization_deserialization_performance

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___SerializationDeserializationWithCompressionPerformance___Json(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer() //vital

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				jsonBytes, err := json.Marshal(x)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(jsonBytes)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				serializedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}

				fooitem := arena.FooItem{}
				err = json.Unmarshal(serializedBytes, &fooitem)
				if err != nil {
					b.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
