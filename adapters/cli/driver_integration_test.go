package cli

import (
	"path/filepath"
	"testing"

	"github.com/bryack/words/adapters/sqlite"
	"github.com/bryack/words/contracts"
)

func TestWordReplacerCLIContract(t *testing.T) {

	t.Run("should replace words in a markdown file", func(t *testing.T) {
		tempDir := t.TempDir()
		input := filepath.Join(tempDir, "input.md")
		output := filepath.Join(tempDir, "output.md")
		provider, err := sqlite.NewSQLiteFormsProvider(sqlite.LoadFromJSONLFile("../../adapters/sqlite/fake.jsonl"))
		if err != nil {
			t.Fatalf("failed to create provider: %v", err)
		}
		driver := &Driver{
			Input:    input,
			Output:   output,
			Old:      "подделка",
			New:      "fake",
			Provider: provider,
		}

		contracts.WordReplacerCLIContract(t, driver)
	})
}
