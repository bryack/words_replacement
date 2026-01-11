package cli

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/bryack/words/internal/replacer"
)

type Driver struct {
	Input  string
	Output string
	Old    string
	New    string
}

type ProductionStubProvider struct{}

func (p ProductionStubProvider) GetForms(word string) (singular, plural []string, err error) {
	if word == "подделка" {
		return []string{"подделка", "подделку"}, []string{"подделки"}, nil
	}
	return nil, nil, fmt.Errorf("word not supported in skeleton")
}

func (d Driver) ReplaceWordsInFile(inputPath, outputPath string) error {
	fsys := os.DirFS(".")
	data, err := fs.ReadFile(fsys, inputPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", inputPath, err)
	}
	provider := ProductionStubProvider{}
	replacer := replacer.NewReplacer(provider)
	repl, err := replacer.Replace(string(data), d.Old, d.New)
	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, []byte(repl), 0644)
}

func (d Driver) ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
