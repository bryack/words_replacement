package cli

import (
	"os"

	"github.com/bryack/words/internal/replacer"
)

type Driver struct {
	Input  string
	Output string
	Old    string
	New    string
}

func (d Driver) ReplaceWordsInFile(inputPath, outputPath string) error {
	fsys := os.DirFS(".")
	data, err := replacer.ReadAndReplace(fsys, inputPath, d.Old, d.New)
	if err != nil {
		return err
	}

	return replacer.WriteFile(outputPath, data)
}

func (d Driver) ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
