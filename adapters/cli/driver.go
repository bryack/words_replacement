package cli

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/bryack/words/adapters/sqlite"
	"github.com/bryack/words/internal/replacer"
)

const DefaultFilePermissions = 0644
const RootDir = "."

type Driver struct {
	Input    string
	Output   string
	Old      string
	New      string
	provider *sqlite.SQLiteFormsProvider
}

func (d *Driver) createReplacer() *replacer.Replacer {
	return replacer.NewReplacer(d.provider)
}

func (d *Driver) ReplaceWordsInFile(inputPath, outputPath string) error {
	fsys := os.DirFS(RootDir)
	data, err := fs.ReadFile(fsys, inputPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", inputPath, err)
	}

	replacer := d.createReplacer()
	repl, err := replacer.Replace(string(data), d.Old, d.New)
	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, []byte(repl), DefaultFilePermissions)
}

func (d *Driver) ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
