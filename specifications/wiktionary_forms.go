package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

// WiktionaryFormsProvider defines essential complexity for word form retrieval
type WiktionaryFormsProvider interface {
	GetForms(word string) (singular, plural []string, err error)
}

// WiktionaryFormsSpecification captures business behavior
func WiktionaryFormsSpecification(t *testing.T, provider WiktionaryFormsProvider) {
	t.Run("should get Russian word forms from Wiktionary", func(t *testing.T) {
		singular, plural, err := provider.GetForms("подделка")
		assert.NoError(t, err)
		assert.True(t, len(singular) > 0, "should return singular forms")
		assert.True(t, len(plural) > 0, "should return plural forms")
		assert.SliceContains(t, singular, "подделка")
		assert.SliceContains(t, plural, "подделки")
	})

	t.Run("should handle non-existent words gracefully", func(t *testing.T) {
		singular, plural, err := provider.GetForms("несуществующееслово")

		assert.Error(t, err)
		assert.True(t, len(singular) == 0, "should be empty")
		assert.True(t, len(plural) == 0, "should be empty")
		assert.Contains(t, err.Error(), "not found")
	})
}
