package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type WordReplacer interface {
	Run(args ...string) error
}

func WordReplacerSpecification(t testing.TB, replacer WordReplacer) {
	t.Helper()

	err := replacer.Run("input.md", "output.md")
	assert.NoError(t, err)
}
