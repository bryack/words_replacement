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

	err := replacer.ReplaceWordsInFile("test_files/input.md", "test_files/output.md")
	assert.NoError(t, err)

	output, err := replacer.ReadFile("test_files/output.md")
	assert.NoError(t, err)
	assert.Contains(t, output, "fakes")
}
