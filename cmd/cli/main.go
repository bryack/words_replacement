package main

import (
	"fmt"
	"os"

	"github.com/bryack/words/adapters/sqlite"
	"github.com/bryack/words/internal/cli"
	"github.com/bryack/words/internal/replacer"
)

const JSONLDataPath = "../../adapters/sqlite/fake.jsonl"
const dbPath = "../../adapters/cli/testdata/testDB.db"

func main() {
	provider, err := sqlite.NewSQLiteFormsProvider(dbPath, sqlite.LoadFromJSONLFile(JSONLDataPath))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer provider.Close()

	wordReplacer := replacer.NewReplacer(provider)

	app := cli.NewCLI(os.Stdin, os.Stdout, wordReplacer)

	rootCmd := cli.NewRootCommand()
	rootCmd.AddCommand(cli.NewReplaceCommand(app))

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
