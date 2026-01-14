package sqlite

import (
	"os/exec"
	"strconv"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestLoadFromJSONL(t *testing.T) {

	t.Run("loads all noun entries from JSONL file", func(t *testing.T) {
		filepath := "fake.jsonl"
		entries, err := LoadFromJSONL(filepath)
		assert.NoError(t, err)

		cmd := exec.Command("wc", "-l", filepath)
		out, err := cmd.Output()
		if err != nil {
			t.Fatalf("failed to analyze %s: %v", filepath, err)
		}
		result := strings.Fields(string(out))
		count, err := strconv.Atoi(result[0])
		assert.NoError(t, err)
		assert.Equal(t, count, len(entries))

		for _, entry := range entries {
			assert.Equal(t, "noun", entry.Pos)
		}
	})
}
