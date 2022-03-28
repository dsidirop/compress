package serialization_deserialization_performance

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/klauspost/compress/arena"
)

var hambaavroNewItem interface{}

func Benchmark___SerializationDeserializationPerformance___HambaAvro(b *testing.B) {
	newitem := arena.ItemSerdeable(nil)
	datasourceArrayLength := len(arena.MainDatasource)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := arena.MainDatasource[i%datasourceArrayLength]

		bytes, err := avro.Marshal(x.HambaAvroSchema, x.Item)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}

		newitem = x.NewEmptyItem()
		err = avro.Unmarshal(x.HambaAvroSchema, bytes, newitem)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	hambaavroNewItem = newitem
}

// jsonbytes1, err := easyjson.Marshal(x.Item)
// jsonbytes2, err := easyjson.Marshal(newitem)
//
// jsonstring1 := string(jsonbytes1)
// jsonstring2 := string(jsonbytes2)
//
// if jsonstring1 != jsonstring2 {
// 	fmt.Println("")
// 	fmt.Println(jsonstring1)
// 	fmt.Println(jsonstring2)
// 	b.Fatalf("Error: %d", i%datasourceArrayLength)
// }
