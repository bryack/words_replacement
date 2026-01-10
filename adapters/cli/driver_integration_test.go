package cli

import (
	"path/filepath"
	"testing"

	"github.com/bryack/words/contracts"
)

func TestWordReplacerCLIContract(t *testing.T) {

	t.Run("should replace words in a markdown file", func(t *testing.T) {
		tempDir := t.TempDir()
		input := filepath.Join(tempDir, "input.md")
		output := filepath.Join(tempDir, "output.md")

		driver := &Driver{
			Input:  input,
			Output: output,
			Old:    "подделка",
			New:    "fake",
		}

		contracts.WordReplacerCLIContract(t, driver)
	})
}
