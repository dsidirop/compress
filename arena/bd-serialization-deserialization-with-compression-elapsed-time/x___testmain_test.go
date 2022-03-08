package serialization_deserialization_performance

import (
	"os"
	"testing"

	"github.com/klauspost/compress/arena"
)

const NUMBER_OF_ITERATIONS = 2_000

func TestMain(m *testing.M) {
	arena.InitTestProvisions()

	exitVal := m.Run()

	os.Exit(exitVal)
}
