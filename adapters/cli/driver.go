package cli

import "fmt"

type Driver struct {
	Input  string
	Output string
}

func (d Driver) ReplaceWordsInFile(inputPath, outputPath string) error {
	return fmt.Errorf("not implemented")
}

func (d Driver) ReadFile(path string) (string, error) {
	return "", nil
}
