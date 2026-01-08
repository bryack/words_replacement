package contracts

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type WordReplacerCLI interface {
	ReplaceWordsInFile(inputPath, outputPath string) error
	ReadFile(path string) (string, error)
}

func WordReplacerCLIContract(t testing.TB, replacer WordReplacerCLI) {
	t.Helper()

	err := replacer.ReplaceWordsInFile("input.md", "output.md")
	assert.NoError(t, err)

	output, err := replacer.ReadFile("output.md")
	assert.NoError(t, err)
	assert.Contains(t, output, "fake")
}
