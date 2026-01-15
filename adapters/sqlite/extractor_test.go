package sqlite

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestExtractAllForms(t *testing.T) {
	dataFromFile, err := os.ReadFile("fake.jsonl")
	assert.NoError(t, err)

	tests := []struct {
		name         string
		data         []byte
		wantSingular []string
		wantPlural   []string
	}{
		{
			name:         "extracts all 6 case forms from real JSONL",
			data:         dataFromFile,
			wantSingular: []string{"подделка", "подделку", "подделки", "подделке", "подделкой", "подделкою"},
			wantPlural:   []string{"подделки", "подделок", "подделкам", "подделками", "подделках"},
		},
		{
			name: "some empty forms and duplicate tags",
			data: []byte(`{
  "word": "подделка",
  "forms": [
	{"form": "", "tags": ["nominative", "singular"]},
	{"form": "", "tags": ["accusative", "singular"]},
	{"form": "", "tags": ["nominative", "plural"]},
	{"form": "подде́лки", "tags": ["accusative", "plural"]},
	{"form": "подде́лки", "tags": ["accusative", "plural"]}
  ]
}`),
			wantSingular: []string{},
			wantPlural:   []string{"подделки"},
		},
		{
			name: "completely empty forms",
			data: []byte(`{
  "word": "подделка",
  "forms": [
  ]
}`),
			wantSingular: []string{},
			wantPlural:   []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var entry KaikkiEntry
			err = json.Unmarshal(tt.data, &entry)
			assert.NoError(t, err)

			s, p := entry.ExtractAllForms()
			assert.Equal(t, tt.wantSingular, s)
			assert.Equal(t, tt.wantPlural, p)
		})
	}
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
