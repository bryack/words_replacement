package acceptance

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Driver struct {
	BinaryPath string
}

func (d *Driver) Replace(text, oldWord, newWord string) (string, error) {
	cmd := exec.Command(d.BinaryPath, oldWord, newWord)

	cmd.Stdin = strings.NewReader(text)

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
