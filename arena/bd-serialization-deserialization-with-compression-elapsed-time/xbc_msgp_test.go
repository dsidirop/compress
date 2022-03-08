package serialization_deserialization_performance

import (
	"bytes"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
	"github.com/tinylib/msgp/msgp"
)

func Test___SerializationDeserializationWithCompressionPerformance___Msgp(t *testing.T) {
	datasource := arena.Datasource
	datasourceArrayLength := len(datasource)

	for _, test := range arena.AllCompressionTestCases {
		t.Run(test.Desc, func(testbed *testing.T) {
			startTime := time.Now()

			for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
				x := datasource[i%datasourceArrayLength]

				serializedBytesBuffer := &bytes.Buffer{}
				err := msgp.Encode(serializedBytesBuffer, &x)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				compressedAndSerializedBytes, err := test.CompressionCallback(serializedBytesBuffer.Bytes())
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				serializedBytes, err := test.DecompressionCallback(compressedAndSerializedBytes)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}

				fooitem := &arena.FooItem{}
				deserializedBytesBuffer := bytes.NewReader(serializedBytes)
				err = msgp.Decode(deserializedBytesBuffer, fooitem)
				if err != nil {
					testbed.Fatalf("Error: %s", err)
				}
			}
			finishTime := time.Now()

			averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

			testbed.Logf("** Msgp %d nanoseconds\n", int64(averageElapsedTime))
		})
	}
}
