package specifications

import (
	"testing"
)

type WordReplacer interface {
	Replace(text, oldWord, newWord string) (string, error)
}

func WordReplacerSpecification(t testing.TB, replacer WordReplacer) {
	tests := []struct {
		name      string
		text      string
		data      string
		oldWord   string
		newWord   string
		want      string
		wantError bool
	}{
		{
			name:    "replaces accusative singular form",
			text:    "Требования к тестам: HTTP-тесты используют ту же подделку",
			oldWord: "подделка",
			newWord: "fake",
			want:    "Требования к тестам: HTTP-тесты используют ту же fake",
		},
		{
			name:    "replaces nominative singular form",
			text:    "Это подделка документа",
			oldWord: "подделка",
			newWord: "fake",
			want:    "Это fake документа",
		},
		{
			name:    "replaces plural forms with 's' suffix",
			text:    "Контракты и подделки для тестирования",
			oldWord: "подделка",
			newWord: "fake",
			want:    "Контракты и fakes для тестирования",
		},
		{
			name:    "replaces multiple forms in same text",
			text:    "Мы нашли подделку и эти подделки нам не нравятся",
			oldWord: "подделка",
			newWord: "fake",
			want:    "Мы нашли fake и эти fakes нам не нравятся",
		},
		{
			name:      "returns error for non-existent word",
			text:      "Some text",
			oldWord:   "несуществующееслово",
			newWord:   "fake",
			wantError: true,
		},
	}
	for _, tt := range tests {
		got, err := replacer.Replace(tt.text, tt.oldWord, tt.newWord)
		if tt.wantError {
			if err == nil {
				t.Errorf("%s: expected error, but got none", tt.name)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: unexpected error: %v", tt.name, err)
		}
		if got != tt.want {
			t.Errorf("%s:\nwant %q\ngot %q", tt.name, tt.want, got)
		}
	}

}
