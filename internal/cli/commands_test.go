package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRootCommand(t *testing.T) {
	cmd := NewRootCommand()
	if cmd.Use != "words" {
		t.Errorf("Expected Use 'words', got %s", cmd.Use)
	}
}

func TestReplaceCommand_Flags(t *testing.T) {
	spyReplacer := &SpyWordReplacer{}
	in := strings.NewReader("")
	out := &bytes.Buffer{}
	cli := NewCLI(in, out, spyReplacer)

	cmd := NewReplaceCommand(cli)

	inputFlag := cmd.Flags().Lookup("input")
	if inputFlag == nil {
		t.Errorf("Expected --input flag to exist")
	}

	dataFlag := cmd.Flags().Lookup("data")
	if dataFlag == nil {
		t.Errorf("Expected --data flag to exist")
	}
}
