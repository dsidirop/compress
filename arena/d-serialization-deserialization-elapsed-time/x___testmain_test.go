package serialization_deserialization_elapsed_time

import (
	"os"
	"testing"

	"github.com/klauspost/compress/arena"
)

const NUMBER_OF_ITERATIONS = 2000

func TestMain(m *testing.M) {
	arena.InitTestProvisions()

	exitVal := m.Run()

	os.Exit(exitVal)
}
