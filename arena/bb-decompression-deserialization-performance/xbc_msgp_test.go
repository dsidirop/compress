package decompression_with_deserialization_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

var msgpNewItem interface{}

func Benchmark___DecompressionAndDeserializationPerformance___Msgp(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		newitem := arena.ItemSerdeable(nil)

		b.Run(test.Desc, func(bench *testing.B) {
			compressedAndSerializedDatasource := [][]byte{} //first serialize and compress

			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				buf := bytes.Buffer{}
				err := msgp.Encode(&buf, x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(buf.Bytes())
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

				serializedBytes, err := test.DecompressionCallback(x)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				newitem = mainItemSpec.NewEmptyItem()
				byteBuffer := bytes.NewBuffer(serializedBytes) //unfortunate necessity

				err = msgp.Decode(byteBuffer, newitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})

		msgpNewItem = newitem
	}
}
