package serialization_deserialization_performance

import (
	"testing"
	"time"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationDeserializationWithCompressionPerformance___HambaAvro(t *testing.T) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		t.Run(test.Desc, func(testbed *testing.T) {

			startTime := time.Now()
			for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
				x := datasource[i%datasourceArrayLength]

				thriftBytes, err := avro.Marshal(arena.Schemas.GoHambaAvro, x)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(thriftBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				decompressedSerializedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				fooitem := &arena.FooItem{}
				err = avro.Unmarshal(arena.Schemas.GoHambaAvro, decompressedSerializedBytes, fooitem)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}
			}
			finishTime := time.Now()

			averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

			testbed.Logf("** HambaAvro %d nanoseconds\n", int64(averageElapsedTime))
		})
	}
}
