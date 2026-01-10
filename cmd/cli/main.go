package main

import (
	"fmt"
	"os"

	"github.com/bryack/words/internal/replacer"
)

func main() {
	provider := replacer.ProductionStubProvider{}
	wordReplacer := replacer.NewReplacer(provider)

	if err := replacer.Run(wordReplacer, os.Stdin, os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
