package decompression_with_deserialization_performance

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/klauspost/compress/arena"
)

func Benchmark___DecompressionAndDeserializationPerformance___ThriftBinary(b *testing.B) {
	ctx := context.TODO()
	datasource := arena.SpecialDatasourcesForIDLMechanisms.Thrift
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			thriftBinarySerializer := thrift.NewTSerializer()     //binary serializer
			thriftBinaryDeserializer := thrift.NewTDeserializer() //binary deserializer

			for i := 0; i < datasourceArrayLength; i++ {
				x := datasource[i]

				serializedBytes, err := thriftBinarySerializer.Write(ctx, x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				x.Bytes = compressedAndSerializedBytes
			}

			bench.ResetTimer() //vital
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength] //and now we deserialize and decompress

				decompressedSerializedBytes, err := test.DecompressionCallback(x.Bytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				emptyitem := x.NewEmptyThriftItem()
				err = thriftBinaryDeserializer.Read(ctx, emptyitem, decompressedSerializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}
		})
	}
}
