package decompression_with_deserialization_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Benchmark___SerializationWithCompressionPerformance___Msgp(b *testing.B) {
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

				buf := bytes.Buffer{}
				err := msgp.Encode(&buf, &x)
				if err != nil {
					panic(err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(buf.Bytes())
				if err != nil {
					panic(err)
				}

				compressedAndSerializedDatasource = append(compressedAndSerializedDatasource, compressedAndSerializedBytes)
			}

			bench.ResetTimer() //vital
			for i := 0; i < bench.N; i++ {
				x := compressedAndSerializedDatasource[i%datasourceArrayLength] //and now we deserialize and decompress

				serializedBytes, err := test.DecompressionCallback(x)
				if err != nil {
					panic(err)
				}

				fooitem := &arena.FooItem{}
				byteBuffer := bytes.NewBuffer(serializedBytes) // unfortunate necessity

				err = msgp.Decode(byteBuffer, fooitem)
				if err != nil {
					panic(err)
				}
			}
		})
	}
}
