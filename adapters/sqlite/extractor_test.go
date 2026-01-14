package sqlite

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestExtractNominativeForms(t *testing.T) {

	t.Run("extracts nominative forms from real JSONL", func(t *testing.T) {
		data, err := os.ReadFile("fake.jsonl")
		assert.NoError(t, err)

		var entry KaikkiEntry
		err = json.Unmarshal(data, &entry)
		assert.NoError(t, err)

		s, p := entry.ExtractNominativeForms()
		assert.Equal(t, "подделка", s)
		assert.Equal(t, "подделки", p)
	})
}

func TestRemoveRussianStress(t *testing.T) {

	t.Run("remove stress from 'подде́лка'", func(t *testing.T) {
		input := "подде́лка"
		got := RemoveRussianStress(input)
		assert.Equal(t, "подделка", got)
	})
	t.Run("remove stress from 'прохо́жий'", func(t *testing.T) {
		input := "прохо́жий"
		got := RemoveRussianStress(input)
		assert.Equal(t, "прохожий", got)
	})
}
