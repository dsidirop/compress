package serialization_deserialization_elapsed_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/klauspost/compress/arena"
)

func Test___SerializationDeserializationElapsedTime___Cbor(t *testing.T) { // https://github.com/fxamacker/cbor
	newitem := arena.ItemSerdeable(nil)
	datasource := arena.MainDatasource
	datasourceArrayLength := len(datasource)

	startTime := time.Now()
	for i := 0; i < NUMBER_OF_ITERATIONS; i++ {
		x := datasource[i%datasourceArrayLength]

		bytes, err := arena.CborStandardEncoder.Marshal(x.Item)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		newitem = x.NewEmptyItem()
		err = arena.CborStandardDecoder.Unmarshal(bytes, newitem)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
	}
	finishTime := time.Now()

	averageElapsedTime := float64(finishTime.Sub(startTime).Nanoseconds()) / NUMBER_OF_ITERATIONS

	fmt.Printf("** CBOR %d nanoseconds\n", int64(averageElapsedTime))
}

// xAsFooItem := x.Item.(*arena.FooItem)
// newitemAsFooItem := newitem.(*arena.FooItem)
//
// micro1 := xAsFooItem.CreatedAt.UnixMicro()
// micro2 := newitemAsFooItem.CreatedAt.UnixMicro()
// if micro1 != micro2 {
// 	  fmt.Println(i)
// }
//
// foo := x.Item.(*arena.SimEventRegisterEventCmd)
// time1, err := time.Parse(time.RFC3339, foo.ApiVersion)
//
//
//
// jsonbytes1, err := json.Marshal(x.Item)
// if err != nil {
// 	panic(err)
// }
//
// jsonbytes2, err := json.Marshal(newitem)
// if err != nil {
// 	panic(err)
// }
//
// jsonstring1 := string(jsonbytes1)
// jsonstring2 := string(jsonbytes2)
// if jsonstring1 != jsonstring2 {
// 	fmt.Println(i)
// }
