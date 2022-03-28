package serialization_deserialization_performance

import (
	"testing"

	"github.com/klauspost/compress/arena"
	"github.com/vmihailenco/msgpack/v5"
)

var messagePackDummyVariable interface{}

func Benchmark___SerializationDeserializationWithCompressionPerformance___MessagePack(b *testing.B) {
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		b.Run(test.Desc, func(bench *testing.B) {
			bench.ResetTimer() //vital

			newitem := arena.ItemSerdeable(nil) //0 always record the result to prevent the compiler eliminating the function call
			for i := 0; i < bench.N; i++ {
				x := datasource[i%datasourceArrayLength]

				bytes, err := msgpack.Marshal(x.Item)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(bytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				serializedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}

				newitem = x.NewEmptyItem()
				err = msgpack.Unmarshal(serializedBytes, &newitem)
				if err != nil {
					bench.Fatalf("Error: %s", err)
				}
			}

			messagePackDummyVariable = newitem //0 always store the result to a package level variable so the compiler cannot eliminate the Benchmark itself
		})
	}

	//0 https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
}
