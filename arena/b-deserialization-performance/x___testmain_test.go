package deserialization_performance

import (
	"os"
	"testing"

	"github.com/klauspost/compress/arena"
)

func TestMain(m *testing.M) {
	arena.InitializeAlternativeDatasourcesFromMainDatasource()

	exitVal := m.Run()

	os.Exit(exitVal)
}
