package serialization_deserialization_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Benchmark___SerializationDeserializationWithCompressionPerformance___Msgp(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer() //vital

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				serializedBytesBuffer := &bytes.Buffer{}
				err := msgp.Encode(serializedBytesBuffer, &x)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytesBuffer.Bytes())
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				decompressedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				fooitem := &arena.FooItem{}
				deserializedBytesBuffer := bytes.NewReader(decompressedBytes)
				err = msgp.Decode(deserializedBytesBuffer, fooitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
