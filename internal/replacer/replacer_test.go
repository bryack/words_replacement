package replacer

import (
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestReplace(t *testing.T) {

	t.Run("one form", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
			want  string
		}{
			{
				name:  "single",
				input: "**Требования к тестам:** HTTP-тесты используют ту же подделку, что и тесты на уровне сервисов",
				want:  "**Требования к тестам:** HTTP-тесты используют ту же fake, что и тесты на уровне сервисов",
			},
			{
				name:  "plural",
				input: "Контракты и подделки для интеграционного тестирования списка дел",
				want:  "Контракты и fakes для интеграционного тестирования списка дел",
			},
			{
				name:  "without",
				input: "Контракты",
				want:  "Контракты",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Replace(tt.input)

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

		data, err := ReadAndReplace(fs, inFile)

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
	})
}
