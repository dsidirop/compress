package serialization_deserialization_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

var msgpDummyVariable interface{}

func Benchmark___SerializationDeserializationWithCompressionPerformance___Msgp(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer() //vital

			newitem := arena.ItemSerdeable(nil)
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				serializedBytesBuffer := &bytes.Buffer{}
				err := msgp.Encode(serializedBytesBuffer, x.Item)
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

				newitem = x.NewEmptyItem()
				err = msgp.Decode(bytes.NewReader(decompressedBytes), newitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}

			msgpDummyVariable = newitem
		})
	}
}
