package decompression_with_deserialization_performance

import (
	"context"
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/klauspost/compress/arena/thfooitem"
)

func Benchmark___DecompressionAndDeserializationPerformance___ThriftCompact(b *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			thriftCompactSerializer := arena.NewThriftCompactSerializer() //compact serializer
			compressedAndSerializedDatasource := [][]byte{}               //first serialize and compress
			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				serializedBytes, err := thriftCompactSerializer.Write(ctx, x)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedDatasource = append(compressedAndSerializedDatasource, compressedAndSerializedBytes)
			}

			bench.ResetTimer()                                                //vital
			thriftCompactDeserializer := arena.NewThriftCompactDeserializer() //compact deserializer
			for i := 0; i < bench.N; i++ {
				x := compressedAndSerializedDatasource[i%datasourceArrayLength] //and now we deserialize and decompress

				decompressedSerializedBytes, err := test.DecompressionCallback(x)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				y := thfooitem.NewTHFooItem()
				err = thriftCompactDeserializer.Read(ctx, y, decompressedSerializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
