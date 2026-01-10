package replacer

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
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
				singular: []string{"подделка", "подделку", "подделки"},
				plural:   []string{"подделки", "подделок", "подделками"},
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

	t.Run("one form", func(t *testing.T) {
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
				old:   "подделка", // Пользователь вводит корень или базовую форму
				new:   "fake",
				want:  "Мы нашли fake и эти fakes нам не нравятся",
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
	})
}

func TestReadAndReplace(t *testing.T) {

	t.Run("read", func(t *testing.T) {
		inFile := "hello.md"

		fs := fstest.MapFS{
			"hello.md": {Data: []byte("**Требования к тестам:** HTTP-тесты используют ту же подделку, что и тесты на уровне сервисов")},
		}
		old := "подделка"
		new := "fake"

		data, err := ReadAndReplace(fs, inFile, old, new)

		if err != nil {
			t.Fatal(err)
		}
		want := "**Требования к тестам:** HTTP-тесты используют ту же fake, что и тесты на уровне сервисов"

		if data != want {
			t.Errorf("want %q, but got %q", want, data)
		}
	})
}

func TestWriteFile(t *testing.T) {

	t.Run("write", func(t *testing.T) {
		tempDir := t.TempDir()
		filename := filepath.Join(tempDir, "hello.md")
		data := "**Требования к тестам:** HTTP-тесты используют ту же подделку, что и тесты на уровне сервисов"

		err := WriteFile(filename, data)

		if err != nil {
			t.Errorf("expected no errors, but got %v", err)
		}

		result, err := os.ReadFile(filename)
		if err != nil {
			t.Errorf("failed to write a file %s: %v", filename, err)
		}

		if string(result) != data {
			t.Errorf("want %q, but got %q", data, string(result))
		}
	})
}
