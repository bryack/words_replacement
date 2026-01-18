package replacer

import (
	"fmt"
	"strings"
	"testing"
)

type StubFormProvider struct {
	data map[string]struct {
		singular []string
		plural   []string
	}
}

func NewStubFormProvider() StubFormProvider {
	return StubFormProvider{
		data: map[string]struct {
			singular []string
			plural   []string
		}{
			"подделка": {
				singular: []string{"подделка", "подделку", "подделки", "подделке", "подделкой", "подделкою"},
				plural:   []string{"подделки", "подделок", "подделкам", "подделками", "подделках"},
			},
		},
	}
}

func (sfp StubFormProvider) GetForms(word string) (singular, plural []string, err error) {
	forms, ok := sfp.data[word]
	if !ok {
		return nil, nil, fmt.Errorf("failed to find %s", word)
	}
	return forms.singular, forms.plural, nil
}

func TestReplace(t *testing.T) {
	stub := NewStubFormProvider()

	tests := []struct {
		name  string
		input string
		old   string
		new   string
		want  string
	}{
		{
			name:  "single",
			input: "**Требования к тестам:** HTTP-тесты используют ту же подделку, что и тесты на уровне сервисов",
			old:   "подделка",
			new:   "fake",
			want:  "**Требования к тестам:** HTTP-тесты используют ту же fake, что и тесты на уровне сервисов",
		},
		{
			name:  "plural",
			input: "Контракты и подделки для интеграционного тестирования подделки",
			old:   "подделка",
			new:   "fake",
			want:  "Контракты и fakes для интеграционного тестирования fakes",
		},
		{
			name:  "without",
			input: "Контракты",
			old:   "подделка",
			new:   "fake",
			want:  "Контракты",
		},
		{
			name:  "inflection_support",
			input: "Мы нашли подделку и эти подделки нам не нравятся",
			old:   "подделка",
			new:   "fake",
			want:  "Мы нашли fake и эти fakes нам не нравятся",
		},
		{
			name:  "longer_form_not_corrupted_by_shorter",
			input: "Он боролся с подделками и другим",
			old:   "подделка",
			new:   "fake",
			want:  "Он боролся с fakes и другим",
		},
		{
			name:  "word_with_exclamation",
			input: "Это подделка!",
			old:   "подделка",
			new:   "fake",
			want:  "Это fake!",
		},
		{
			name:  "word_with_period",
			input: "Работа с подделками.",
			old:   "подделка",
			new:   "fake",
			want:  "Работа с fakes.",
		},
		{
			name:  "word_with_comma",
			input: "Вот подделка, которую нашли",
			old:   "подделка",
			new:   "fake",
			want:  "Вот fake, которую нашли",
		},
		{
			name:  "should_not_replace_inside_word",
			input: "Это неподделка",
			old:   "подделка",
			new:   "fake",
			want:  "Это неподделка",
		},
		{
			name:  "word with upper-case letter",
			input: "Это ПодДелКа",
			old:   "подделка",
			new:   "fake",
			want:  "Это fake",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replacer := NewReplacer(stub)
			got, err := replacer.Replace(tt.input, tt.old, tt.new)
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Errorf("want %q, but got %q", tt.want, got)
			}
		})
	}
}

func TestReplaceWord(t *testing.T) {
	tests := []struct {
		name  string
		input string
		old   string
		new   string
		want  string
	}{
		{
			name:  "replaces word surrounded by spaces",
			input: "это подделка точно",
			old:   "подделка",
			new:   "fake",
			want:  "это fake точно",
		},
		{
			name:  "replace several matches",
			input: "это подделка точно, (подделка).",
			old:   "подделка",
			new:   "fake",
			want:  "это fake точно, (fake).",
		},
		{
			name:  "does not replace inside word",
			input: "это неподделка точно, вон та - подделка-то. подделкаподделка",
			old:   "подделка",
			new:   "fake",
			want:  "это неподделка точно, вон та - fake-то. подделкаподделка",
		},
		{
			name:  "should replace with capital letters",
			input: "это Подделка_ точно, вон та - поДдеЛка.",
			old:   "подделка",
			new:   "fake",
			want:  "это fake_ точно, вон та - fake.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replaceRussianWord(tt.input, tt.old, tt.new)
			if got != tt.want {
				t.Errorf("want %q, but got %q", tt.want, got)
			}
		})
	}
}

func BenchmarkReplaceRussianWord(b *testing.B) {
	text := "Это подделка точно, вон та - подделка-то. Мы нашли подделку и эти подделки нам не нравятся."
	old := "подделка"
	new := "fake"

	for b.Loop() {
		replaceRussianWord(text, old, new)
	}
}

func BenchmarkReplaceRussianWordNoMatches(b *testing.B) {
	text := "Это точно фальшивка, вон та — фальшивка-то. Мы нашли фальшивку и эти фальшивки нам не нравятся."
	old := "подделка"
	new := "fake"

	for b.Loop() {
		replaceRussianWord(text, old, new)
	}
}

func BenchmarkReplaceRussianWordManyMatches(b *testing.B) {
	text := strings.Repeat("подделка", 100)
	old := "подделка"
	new := "fake"
	b.ResetTimer()

	for b.Loop() {
		replaceRussianWord(text, old, new)
	}
}

func BenchmarkReplace(b *testing.B) {
	stub := NewStubFormProvider()
	replacer := NewReplacer(stub)
	input := "Это подделка точно, вон та - подделка-то. Мы нашли подделку и эти подделки нам не нравятся."

	for b.Loop() {
		replacer.Replace(input, "подделка", "fake")
	}
}
