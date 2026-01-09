package acceptance

import (
	"bytes"
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

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}
