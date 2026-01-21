// Package cli provides a test adapter for the WordReplacerCLI contract.
//
// This package is designed for integration testing and implements the
// contracts.WordReplacerCLI interface. It provides a file-based driver
// that reads input files, performs word replacement using a FormsProvider,
// and writes output files for verification in tests.
//
// This adapter is not intended for production use. For production CLI
// functionality, see internal/cli package.

package cli

import (
	"fmt"
	"os"

	"github.com/bryack/words/internal/replacer"
)

const DefaultFilePermissions = 0644
const RootDir = "."

// Driver is a test adapter that implements the contracts.WordReplacerCLI interface.
// It performs file-based word replacement for integration testing.
type Driver struct {
	Input          string
	Output         string
	Old            string
	New            string
	ExpectedOutput string
	Provider       replacer.FormsProvider
}

func (d *Driver) createReplacer() *replacer.Replacer {
	return replacer.NewReplacer(d.Provider)
}

// ReplaceWordsInFile reads the input file, performs word replacement,
// and writes the result to the output file.
func (d *Driver) ReplaceWordsInFile() error {
	data, err := os.ReadFile(d.Input)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", d.Input, err)
	}

	replacer := d.createReplacer()
	repl, err := replacer.Replace(string(data), d.Old, d.New)
	if err != nil {
		return err
	}

	return os.WriteFile(d.Output, []byte(repl), DefaultFilePermissions)
}

// ReadOutput reads and returns the content of the output file.
func (d *Driver) ReadOutput() (string, error) {
	data, err := os.ReadFile(d.Output)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadExpectedOutput reads and returns the content of the expected output file.
func (d *Driver) ReadExpectedOutput() (string, error) {
	data, err := os.ReadFile(d.ExpectedOutput)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
