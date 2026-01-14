package main_test

import (
	"testing"

	"github.com/bryack/words/adapters/wiktionary"
	"github.com/bryack/words/specifications"
)

func TestWiktionaryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping external API test")
	}

	t.Run("should integrate with real Wiktionary API", func(t *testing.T) {
		provider, err := wiktionary.NewProvider("https://ru.wiktionary.org/w/api.php")
		if err != nil {
			t.Error(err)
		}
		specifications.WiktionaryFormsSpecification(t, provider)
	})
}
