package main_test

import (
	"path/filepath"
	"testing"

	"github.com/bryack/words/adapters/cli"
	"github.com/bryack/words/specifications"
)

func TestWordReplacementSpecification(t *testing.T) {

	t.Run("should replace words in a markdown file", func(t *testing.T) {
		tempDir := t.TempDir()
		input := filepath.Join(tempDir, "input.md")
		output := filepath.Join(tempDir, "output.md")

		driver := &cli.Driver{
			Input:  input,
			Output: output,
		}

		specifications.WordReplacerSpecification(t, driver)
	})
}
