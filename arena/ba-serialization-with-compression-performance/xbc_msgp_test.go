package serialization_with_compression_performance

import (
	"bytes"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

var msgpBytes []byte

func Benchmark___SerializationAndCompressionPerformance___Msgp(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	var results []byte
	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			//b.ResetTimer() //something goes wrong when doing this in this particular test    weird

			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				buf := bytes.Buffer{}
				err := msgp.Encode(&buf, x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				results, err = test.CompressionCallback(buf.Bytes())
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})

		msgpBytes = results
	}
}
