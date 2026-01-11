package main_test

import (
	"testing"

	"github.com/bryack/words/adapters/wiktionary"
	"github.com/bryack/words/specifications"
)

func TestWiktionaryIntegration(t *testing.T) {

	t.Run("should integrate with real Wiktionary API", func(t *testing.T) {
		provider := wiktionary.NewProvider("https://ru.wiktionary.org/w/api.php")
		specifications.WiktionaryFormsSpecification(t, provider)
	})
}
