package arena

import (
	"encoding/json"
	"testing"

	"github.com/klauspost/compress/arena"
)

func Benchmark___Deserialization___Json(t *testing.B) {
	item := arena.FooItem{}
	datasourceArrayLength := len(arena.Datasource)

	for i := 0; i < t.N; i++ {
		bytes := arena.DatasourceSerializedJson[i%datasourceArrayLength]

		err := json.Unmarshal(bytes, &item)
		if err != nil {
			panic(err)
		}
	}
}
