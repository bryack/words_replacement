package acceptance

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Driver struct {
	BinaryPath string
	DataFile   string
	TempDir    string
}

func (d *Driver) Replace(text, oldWord, newWord string) (string, error) {
	inputFile := filepath.Join(d.TempDir, "input.txt")
	if err := os.WriteFile(inputFile, []byte(text), 0644); err != nil {
		return "", fmt.Errorf("failed to create file %s: %w", inputFile, err)
	}
	args := []string{
		"replace",
		"--input", inputFile,
		"--old", oldWord,
		"--new", newWord,
	}

	if d.DataFile != "" {
		args = append(args, "--data", d.DataFile)
	}

	cmd := exec.Command(d.BinaryPath, args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Go run Error: \n%q\n", stderr.String())
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}
