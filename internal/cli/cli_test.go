package cli

import (
	"bytes"
	"strings"
	"testing"
)

type SpyWordReplacer struct {
	called                          bool
	gotText, gotOld, gotNew, result string
	err                             error
}

func (swr *SpyWordReplacer) Replace(text, oldWord, newWord string) (string, error) {
	swr.called = true
	swr.gotText = text
	swr.gotOld = oldWord
	swr.gotNew = newWord
	return swr.result, swr.err
}

func TestHello(t *testing.T) {
	t.Run("TestCase", func(t *testing.T) {
		input := "Требования к тестам: HTTP-тесты используют ту же подделку"
		want := "Требования к тестам: HTTP-тесты используют ту же fake"
		spyReplacer := &SpyWordReplacer{
			result: want,
		}
		in := strings.NewReader(input)
		out := &bytes.Buffer{}
		app := NewCLI(in, out, spyReplacer)

		err := app.Run([]string{
			"подделка",
			"fake",
		})
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}

		got := out.String()

		if got != want {
			t.Errorf("want %q, but got %q", want, got)
		}

		if spyReplacer.gotOld != "подделка" || spyReplacer.gotNew != "fake" {
			t.Errorf("Replacer получил неверные аргументы: old=%q, new=%q",
				spyReplacer.gotOld, spyReplacer.gotNew)
		}

		if spyReplacer.gotText != input {
			t.Errorf("Spy got incorrect text %q, expected: %q", spyReplacer.gotText, input)
		}
	})
}

func TestCLI_RunWithFiles(t *testing.T) {

}
