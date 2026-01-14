package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type WordReplacer interface {
	Replace(text, oldWord, newWord string) (string, error)
}

func WordReplacerSpecification(t testing.TB, replacer WordReplacer) {
	text := "Требования к тестам: HTTP-тесты используют ту же подделка"
	oldWord := "подделка"
	newWord := "fake"

	got, err := replacer.Replace(text, oldWord, newWord)

	assert.NoError(t, err)
	assert.Equal(t, "Требования к тестам: HTTP-тесты используют ту же fake", got)
}
