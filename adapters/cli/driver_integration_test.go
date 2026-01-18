package cli

import (
	"testing"

	"github.com/bryack/words/adapters/sqlite"
	"github.com/bryack/words/contracts"
)

func TestWordReplacerCLIContract(t *testing.T) {

	t.Run("should replace words in a markdown file", func(t *testing.T) {
		input := "testdata/input.md"
		output := "testdata/output.md"
		expected := "testdata/expected_output.md"
		dbPath := "testdata/testDB.db"
		provider, err := sqlite.NewSQLiteFormsProvider(dbPath, sqlite.LoadFromJSONLFile("../../adapters/sqlite/fake.jsonl"))
		if err != nil {
			t.Fatalf("failed to create provider: %v", err)
		}

		driver := &Driver{
			Input:          input,
			Output:         output,
			Old:            "подделка",
			New:            "fake",
			ExpectedOutput: expected,
			Provider:       provider,
		}

		contracts.WordReplacerCLIContract(t, driver)
	})
}
