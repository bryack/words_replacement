package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bryack/words/adapters/cli"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("path of files is required")
	}

	driver := cli.Driver{
		Input:  os.Args[1],
		Output: os.Args[2],
	}

	if err := driver.ReplaceWordsInFile(driver.Input, driver.Output); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
