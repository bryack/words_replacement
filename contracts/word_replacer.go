package contracts

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type WordReplacerCLI interface {
	ReplaceWordsInFile() error
	ReadOutput() (string, error)
}

func WordReplacerCLIContract(t testing.TB, replacer WordReplacerCLI) {
	t.Helper()

	err := replacer.ReplaceWordsInFile()
	assert.NoError(t, err)

	output, err := replacer.ReadOutput()
	assert.NoError(t, err)
	assert.Contains(t, output, "fake")
}
