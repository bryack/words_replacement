package acceptance

import (
	"os/exec"
)

type Driver struct {
	Input  string
	Output string
}

func NewDriver(input, output string) *Driver {

	return &Driver{
		Input:  input,
		Output: output,
	}
}

func (d *Driver) Run(args ...string) error {
	cmd := exec.Command("go", "run", "./cmd/cli", d.Input, d.Output)
	cmd.Dir = "../../"
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
