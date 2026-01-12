package cli

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/bryack/words/adapters/wiktionary"
	"github.com/bryack/words/internal/replacer"
)

const DefaultFilePermissions = 0644
const SupportedWord = "подделка"
const RootDir = "."

type Driver struct {
	Input  string
	Output string
	Old    string
	New    string
}

func (d Driver) createReplacer() (*replacer.Replacer, error) {
	provider, err := wiktionary.NewProvider("https://en.wiktionary.org/w/api.php")
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %w", err)
	}
	return replacer.NewReplacer(provider), nil
}

func (d Driver) ReplaceWordsInFile(inputPath, outputPath string) error {
	fsys := os.DirFS(RootDir)
	data, err := fs.ReadFile(fsys, inputPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", inputPath, err)
	}

	replacer, err := d.createReplacer()
	if err != nil {
		return fmt.Errorf("failed to create replacer: %w", err)
	}
	repl, err := replacer.Replace(string(data), d.Old, d.New)
	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, []byte(repl), DefaultFilePermissions)
}

func (d Driver) ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
