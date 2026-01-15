package wiktionary

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/bryack/words/specifications"
)

type ParseAdapter struct {
	parser *RussianNounParser
}

func (pa *ParseAdapter) GetForms(word string) (singular, plural []string, err error) {

	wikitext := "{{ru-noun|подделка|f|подделки|подделок}}"
	if word != "подделка" {
		wikitext = ""
	}

	return pa.parser.ParseForms(wikitext)
}

func TestRussianNounParser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping external API test")
	}
	t.Run("specification test", func(t *testing.T) {
		russianNounParser := NewRussianNounParser()
		parser := &ParseAdapter{
			parser: russianNounParser,
		}
		specifications.WiktionaryFormsSpecification(t, parser)
	})
}

func TestParseForms(t *testing.T) {
	tests := []struct {
		name         string
		wikitext     string
		wantSingular []string
		wantPlural   []string
		wantError    error
	}{
		{
			name:         "feminine noun подделка",
			wikitext:     "{{ru-noun|подделка|f|подделки|подделок}}",
			wantSingular: []string{"подделка"},
			wantPlural:   []string{"подделки", "подделок"},
			wantError:    nil,
		},
		{
			name:         "masculine noun стол",
			wikitext:     "{{ru-noun|стол|m|столы|столов}}",
			wantSingular: []string{"стол"},
			wantPlural:   []string{"столы", "столов"},
			wantError:    nil,
		},
		{
			name:         "malformed template",
			wikitext:     "{{ru-noun|incomplete",
			wantSingular: nil,
			wantPlural:   nil,
			wantError:    fmt.Errorf("not found"),
		},
		{
			name:         "template with empty parameters",
			wikitext:     "{{ru-noun|подделка|f|подделки|}}", // Empty params
			wantSingular: nil,
			wantPlural:   nil,
			wantError:    fmt.Errorf("empty parameter"),
		},
		{
			name:         "non-Russian template",
			wikitext:     "{{en-noun|cat|cats}}",
			wantSingular: nil,
			wantPlural:   nil,
			wantError:    fmt.Errorf("not found"), // Current: "not found"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewRussianNounParser()
			s, p, err := parser.ParseForms(tt.wikitext)

			assert.Equal(t, tt.wantSingular, s)
			assert.Equal(t, tt.wantPlural, p)
			assert.Equal(t, tt.wantError, err)
		})
	}
}
