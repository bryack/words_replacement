package replacer

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/bryack/words/specifications"
)

type FormsProvider interface {
	GetForms(word string) (singular, plural []string, err error)
}

type Replacer struct {
	provider FormsProvider
}

func NewReplacer(fp FormsProvider) *Replacer {
	return &Replacer{provider: fp}
}

type ProductionStubProvider struct{}

func (p ProductionStubProvider) GetForms(word string) (singular, plural []string, err error) {
	if word == "подделка" {
		return []string{"подделка", "подделку"}, []string{"подделки"}, nil
	}
	return nil, nil, fmt.Errorf("word not supported in skeleton")
}

func Run(r specifications.WordReplacer, in io.Reader, out io.Writer, args []string) error {
	var content []byte
	var err error

	content, err = io.ReadAll(in)

	if err != nil {
		return err
	}

	result, err := r.Replace(string(content), args[0], args[1])
	if err != nil {
		return err
	}

	fmt.Fprint(out, result)
	return nil
}

func (r *Replacer) Replace(input, old, new string) (string, error) {
	sing, plur, err := r.provider.GetForms(old)
	if err != nil {
		return "", fmt.Errorf("failed to get forms of %s: %w", old, err)
	}
	result := input

	for _, form := range plur {
		result = strings.ReplaceAll(result, form, new+"s")
	}

	for _, form := range sing {
		result = strings.ReplaceAll(result, form, new)
	}
	return result, err
}

func ReadAndReplace(fsys fs.FS, filename, old, new string) (string, error) {
	data, err := fs.ReadFile(fsys, filename)
	if err != nil {
		return "", err
	}
	provider := ProductionStubProvider{}
	wordReplacer := NewReplacer(provider)
	repl, err := wordReplacer.Replace(string(data), old, new)
	if err != nil {
		return "", err
	}

	return repl, nil
}

func WriteFile(filename, data string) error {
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}
