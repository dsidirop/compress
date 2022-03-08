package serialization_deserialization_performance

import (
	"testing"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/arena"
)

func Test___SerializationDeserializationWithCompressionPerformance___Cbor(t *testing.T) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		t.Run(test.Desc, func(testbed *testing.T) {
			startTime := time.Now()

			for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
				x := datasource[i%datasourceArrayLength]

				serializedBytes, err := cbor.Marshal(x)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				decompressedSerializedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				fooitem := &arena.FooItem{}
				err = cbor.Unmarshal(decompressedSerializedBytes, fooitem)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}
			}
			finishTime := time.Now()

			averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

			testbed.Logf("** Cbor %d nanoseconds\n", int64(averageElapsedTime))
		})
	}
}
