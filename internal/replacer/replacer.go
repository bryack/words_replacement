package replacer

import (
	"fmt"
	"os"
	"strings"
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

func WriteFile(filename, data string) error {
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}
