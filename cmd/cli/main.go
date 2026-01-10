package main

import (
	"fmt"
	"os"

	"github.com/bryack/words/internal/replacer"
)

func main() {
	if err := replacer.Run(os.Stdin, os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
