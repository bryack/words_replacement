package sqlite

import (
	"slices"
	"strings"
)

const russianStressMark = '\u0301'

// ExtractTwoForms extracts nominative and accusative singular and plural forms from the entry.
// Stress marks are removed from the returned forms.
// Returns empty string slices if the forms are not found.
func (entry *KaikkiEntry) ExtractTwoForms() (singular, plural []string) {
	for _, form := range entry.Forms {
		if len(form.Form) == 0 {
			continue
		}
		// Nominative
		if containsAll(form.Tags, []string{"nominative", "singular"}) {
			singular = append(singular, RemoveRussianStress(form.Form))
		}

		if containsAll(form.Tags, []string{"nominative", "plural"}) {
			plural = append(plural, RemoveRussianStress(form.Form))
		}

		// Accusative
		if containsAll(form.Tags, []string{"accusative", "singular"}) {
			singular = append(singular, RemoveRussianStress(form.Form))
		}
		if containsAll(form.Tags, []string{"accusative", "plural"}) {
			plural = append(plural, RemoveRussianStress(form.Form))
		}
	}
	return deduplicate(singular), deduplicate(plural)
}

func deduplicate(forms []string) []string {
	if len(forms) == 0 {
		return []string{}
	}
	seen := make(map[string]bool, len(forms))
	result := make([]string, 0, len(forms))
	for _, form := range forms {
		if !seen[form] {
			seen[form] = true
			result = append(result, form)
		}
	}
	return result
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
