package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type WordReplacer interface {
	ReplaceWordsInFile(inputPath, outputPath string) error
	ReadFile(path string) (string, error)
}

func WordReplacerSpecification(t testing.TB, replacer WordReplacer) {
	t.Helper()

	err := replacer.ReplaceWordsInFile("input.md", "output.md")
	assert.NoError(t, err)

	output, err := replacer.ReadFile("output.md")
	assert.NoError(t, err)
	assert.Contains(t, output, "fake")
}
