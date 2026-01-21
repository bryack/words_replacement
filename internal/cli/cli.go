package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/bryack/words/specifications"
)

const RequiredArgs = 2

type CLI struct {
	in  io.Reader
	out io.Writer
	r   specifications.WordReplacer
}

func NewCLI(in io.Reader, out io.Writer, r specifications.WordReplacer) *CLI {
	return &CLI{in: in, out: out, r: r}
}

func (cli *CLI) Run(args []string) error {
	if len(args) < RequiredArgs {
		return fmt.Errorf("usage: <old_word> <new_word>")
	}

	content, err := io.ReadAll(cli.in)

	if err != nil {
		return err
	}

	result, err := cli.r.Replace(string(content), args[0], args[1])
	if err != nil {
		return err
	}

	fmt.Fprint(cli.out, result)
	return nil
}

func (cli *CLI) RunWithFiles(inputFile, oldWord, newWord string) error {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	result, err := cli.r.Replace(string(content), oldWord, newWord)
	if err != nil {
		return err
	}

	fmt.Fprint(cli.out, result)
	return nil
}
