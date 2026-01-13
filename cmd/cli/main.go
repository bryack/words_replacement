package main

import (
	"fmt"
	"os"

	adapters "github.com/bryack/words/adapters/cli"
	"github.com/bryack/words/internal/cli"
	"github.com/bryack/words/internal/replacer"
)

func main() {
	provider := adapters.ProductionStubProvider{}
	wordReplacer := replacer.NewReplacer(provider)

	app := cli.NewCLI(os.Stdin, os.Stdout, wordReplacer)

	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
