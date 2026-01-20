package testhelpers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestFileHelpers(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		helper   func(*testing.T, string)
	}{
		{
			name:     "creates JSONL file with Russian test data",
			filename: "test.jsonl",
			helper:   CreateTestJSONLFile,
		},
		{
			name:     "creates input file with Russian test content",
			filename: "test.md",
			helper:   CreateTestFiles,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			filepath := filepath.Join(tempDir, tt.filename)
			tt.helper(t, filepath)

			got, err := os.ReadFile(filepath)
			assert.NoError(t, err)
			assert.True(t, len(got) > 0, "file should not be empty")

			content := string(got)
			assert.Contains(t, content, "подделка", "should contain Russian test word")
		})
	}
}
