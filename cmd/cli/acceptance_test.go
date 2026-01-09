package main_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/bryack/words/adapters/acceptance"
	"github.com/bryack/words/specifications"
)

func TestWordReplacerSpecification(t *testing.T) {

	t.Run("should replace words in a markdown file", func(t *testing.T) {
		tempDir := t.TempDir()
		input := filepath.Join(tempDir, "input")
		output := filepath.Join(tempDir, "output")

		err := os.WriteFile(input, []byte("Требования к тестам: HTTP-тесты используют ту же подделку"), 0644)
		assert.NoError(t, err)

		driver := acceptance.NewDriver(input, output)
		specifications.WordReplacerSpecification(t, driver)

		got, err := os.ReadFile(output)
		assert.NoError(t, err)
		assert.Equal(t, "Требования к тестам: HTTP-тесты используют ту же fake", string(got))
	})
}
