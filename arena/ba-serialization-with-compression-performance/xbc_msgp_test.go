package serialization_with_compression_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Benchmark___SerializationAndCompressionPerformance___Msgp(b *testing.B) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			//b.ResetTimer() //something goes wrong when doing this in this particular test    weird

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				buf := bytes.Buffer{}
				err := msgp.Encode(&buf, &x)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				test.CompressionCallback(buf.Bytes())
			}
		})
	}
}
