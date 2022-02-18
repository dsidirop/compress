package serialization

import (
	"log"
	"testing"
	"unsafe"

	"github.com/klauspost/compress/arena"
)

func Test___SerializationMessageSizeFootprint___RawGolang(t *testing.T) {
	item := arena.Datasource[0]

	log.Printf("** RawGolang %d bytes\n", unsafe.Sizeof(item))
}
