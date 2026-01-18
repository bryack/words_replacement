package replacer

import (
	"fmt"
	"strings"
)

const pluralSuffix = "s"

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
		result = replaceRussianWord(result, form, new+pluralSuffix)
	}

	for _, form := range sing {
		result = replaceRussianWord(result, form, new)
	}
	return result, nil
}

func replaceRussianWord(text, old, new string) string {
	runeText := []rune(text)
	runeOld := []rune(old)

	var builder strings.Builder
	builder.Grow(len(text))
	i := 0
	for i < len(runeText) {
		if i <= len(runeText)-len(runeOld) && matchesAt(runeText, i, runeOld) {
			leftOK := i == 0 || !isCyrillic(runeText[i-1])
			rightOK := i+len(runeOld) >= len(runeText) || !isCyrillic(runeText[i+len(runeOld)])
			if leftOK && rightOK {
				builder.WriteString(new)
				i += len(runeOld)
				continue
			}
		}
		builder.WriteRune(runeText[i])
		i++
	}
	return builder.String()
}

func matchesAt(runes []rune, pos int, pattern []rune) bool {
	if pos+len(pattern) > len(runes) {
		return false
	}

	for j := 0; j < len(pattern); j++ {
		if !equalFoldRune(runes[pos+j], pattern[j]) {
			return false
		}
	}
	return true
}

func isCyrillic(r rune) bool {
	return (r >= 'а' && r <= 'я') || (r >= 'А' && r <= 'Я') || r == 'ё' || r == 'Ё'
}

func equalFoldRune(rune1, rune2 rune) bool {
	return strings.EqualFold(string(rune1), string(rune2))
}
