package replacer

import (
	"fmt"
	"sort"
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

	sortSlices(sing)
	sortSlices(plur)

	result := input

	for _, form := range plur {
		result = strings.ReplaceAll(result, form, new+"s")
	}

	for _, form := range sing {
		result = strings.ReplaceAll(result, form, new)
	}
	return result, err
}

func sortSlices(slice []string) {
	sort.Slice(slice, func(i, j int) bool {
		return len(slice[i]) > len(slice[j])
	})
}
