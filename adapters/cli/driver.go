package cli

import (
	"fmt"
	"os"

	"github.com/bryack/words/internal/replacer"
)

const DefaultFilePermissions = 0644
const RootDir = "."

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

func (d *Driver) ReadOutput() (string, error) {
	data, err := os.ReadFile(d.Output)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (d *Driver) ReadExpectedOutput() (string, error) {
	data, err := os.ReadFile(d.ExpectedOutput)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
