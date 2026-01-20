package cli

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/bryack/words/testhelpers"
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

func TestCLI_Run(t *testing.T) {
	t.Run("replaces words using injected replacer with stdin input", func(t *testing.T) {
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

	tempDir := t.TempDir()
	inputFile := filepath.Join(tempDir, "test.md")
	dataFile := filepath.Join(tempDir, "data.jsonl")

	testhelpers.CreateTestFiles(t, inputFile)
	testhelpers.CreateTestJSONLFile(t, dataFile)

	spyReplacer := &SpyWordReplacer{
		result: "expected result",
	}
	in := strings.NewReader("")
	out := &bytes.Buffer{}
	cli := NewCLI(in, out, spyReplacer)

	err := cli.RunWithFiles(inputFile, dataFile, "подделка", "fake")
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	assert.Equal(t, "expected result", out.String())
	assert.True(t, spyReplacer.called)
	assert.Equal(t, "подделка", spyReplacer.gotOld)
	assert.Equal(t, "fake", spyReplacer.gotNew)
	assert.Contains(t, spyReplacer.gotText, "подделка")
}
