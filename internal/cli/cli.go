package cli

import (
	"fmt"
	"io"

	"github.com/bryack/words/specifications"
)

type CLI struct {
	in  io.Reader
	out io.Writer
	r   specifications.WordReplacer
}

func NewCLI(in io.Reader, out io.Writer, r specifications.WordReplacer) *CLI {
	return &CLI{in: in, out: out, r: r}
}

func (c *CLI) Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: <old_word> <new_word>")
	}

	content, err := io.ReadAll(c.in)

	if err != nil {
		return err
	}

	result, err := c.r.Replace(string(content), args[0], args[1])
	if err != nil {
		return err
	}

	fmt.Fprint(c.out, result)
	return nil
}
