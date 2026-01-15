package sqlite

import (
	"slices"
	"strings"
)

const russianStressMark = '\u0301'

// ExtractAllForms extracts all 6 Russian case forms:
// nominative, accusative, genitive, dative, instrumental, prepositional.
// Stress marks are removed from the returned forms.
// Returns empty string slices if the forms are not found.
func (entry *KaikkiEntry) ExtractAllForms() (singular, plural []string) {
	caseNames := []string{"nominative",
		"accusative",
		"genitive",
		"dative",
		"instrumental",
		"prepositional",
	}
	for _, caseName := range caseNames {
		s, p := entry.extractCases(caseName)
		singular = append(singular, s...)
		plural = append(plural, p...)
	}

	return deduplicate(singular), deduplicate(plural)
}

func (entry *KaikkiEntry) extractCases(caseName string) (singular, plural []string) {
	for _, form := range entry.Forms {
		if len(form.Form) == 0 {
			continue
		}
		if containsAll(form.Tags, []string{caseName, "singular"}) {
			singular = append(singular, RemoveRussianStress(form.Form))
		}
		if containsAll(form.Tags, []string{caseName, "plural"}) {
			plural = append(plural, RemoveRussianStress(form.Form))
		}
	}
	return singular, plural
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
