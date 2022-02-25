package serialization_performance

import (
	"os"
	"testing"

	"github.com/klauspost/compress/arena"
)

func TestMain(m *testing.M) {
	arena.InitTestProvisions()

	exitVal := m.Run()

	os.Exit(exitVal)
}
