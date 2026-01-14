package sqlite

import (
	"slices"
	"strings"
)

const russianStressMark = '\u0301'

// ExtractNominativeForms extracts nominative singular and plural forms from the entry.
// Stress marks are removed from the returned forms.
// Returns empty strings if the forms are not found.
func (entry *KaikkiEntry) ExtractNominativeForms() (singular, plural string) {
	for _, form := range entry.Forms {
		if containsAll(form.Tags, []string{"nominative", "singular"}) {
			singular = RemoveRussianStress(form.Form)
		}
		if containsAll(form.Tags, []string{"nominative", "plural"}) {
			plural = RemoveRussianStress(form.Form)
		}
	}
	return singular, plural
}

// containsAll checks if slice contains all items.
func containsAll(tags, items []string) bool {
	for _, item := range items {
		if !slices.Contains(tags, item) {
			return false
		}
	}
	return true
}

// RemoveRussianStress removes Unicode combining acute accent (U+0301) used for stress marking in Russian text.
func RemoveRussianStress(word string) string {
	return strings.Map(func(r rune) rune {
		if r == russianStressMark {
			return -1
		}
		return r
	}, word)
}
