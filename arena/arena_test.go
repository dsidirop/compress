package arena

import (
	"fmt"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

func BenchmarkTestRecursiveFibonacci_10(t *testing.B) {
	type Item struct {
		Foo string
	}

	b, err := msgpack.Marshal(&Item{Foo: "bar"})
	if err != nil {
		panic(err)
	}

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item.Foo)
	// Output: bar
}
