package main

import (
	"os"

	"github.com/bryack/words/internal/cli"
)

func main() {
	rootCmd := cli.NewRootCommand()
	rootCmd.AddCommand(cli.NewReplaceCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
