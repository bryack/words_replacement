package main

import (
	"fmt"
	"os"

	"github.com/bryack/words/internal/replacer"
)

func main() {
	fsys := os.DirFS(".")
	if err := replacer.Run(fsys, os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
