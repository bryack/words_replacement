package cli

import (
	"testing"
)

func TestRootCommand(t *testing.T) {
	cmd := NewRootCommand()
	if cmd.Use != "words" {
		t.Errorf("Expected Use 'words', got %s", cmd.Use)
	}
}

func TestReplaceCommand_Flags(t *testing.T) {
	cmd := NewReplaceCommand()

	inputFlag := cmd.Flags().Lookup("input")
	if inputFlag == nil {
		t.Errorf("Expected --input flag to exist")
	}

	dataFlag := cmd.Flags().Lookup("data")
	if dataFlag == nil {
		t.Errorf("Expected --data flag to exist")
	}
}
