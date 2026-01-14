package main

import (
	"fmt"
	"os"

	"github.com/bryack/words/adapters/sqlite"
	"github.com/bryack/words/internal/cli"
	"github.com/bryack/words/internal/replacer"
)

const JSONLDataPath = "../../adapters/sqlite/fake.jsonl"

func main() {
	provider, err := sqlite.NewSQLiteFormsProvider(sqlite.LoadFromJSONLFile(JSONLDataPath))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	wordReplacer := replacer.NewReplacer(provider)

	app := cli.NewCLI(os.Stdin, os.Stdout, wordReplacer)

	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
