package main

import (
	"fmt"
	"os"

	"github.com/bryack/words/adapters/wiktionary"
	"github.com/bryack/words/internal/cli"
	"github.com/bryack/words/internal/replacer"
)

func main() {
	provider, err := wiktionary.NewProvider("https://en.wiktionary.org/w/api.php")
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
