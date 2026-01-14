package sqlite

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestKaikkiEntry(t *testing.T) {

	t.Run("unmarshaling real JSONL", func(t *testing.T) {
		data, err := os.ReadFile("fake.jsonl")
		assert.NoError(t, err)

		var entry KaikkiEntry
		err = json.Unmarshal(data, &entry)
		assert.NoError(t, err)
		assert.Equal(t, "подделка", entry.Word)
		assert.Equal(t, "noun", entry.Pos)

		assert.True(t, len(entry.Forms) > 0, "forms should not be empty")

		firstForm := entry.Forms[0]
		assert.NotEqual(t, "", firstForm.Form, "form should nor be empty")
		assert.True(t, len(firstForm.Tags) > 0, "tags array should not be empty")
	})

	t.Run("empty forms", func(t *testing.T) {
		jsonData := `{"word": "test", "pos": "noun", "forms": []}`
		var entry KaikkiEntry
		err := json.Unmarshal([]byte(jsonData), &entry)
		assert.NoError(t, err)
		assert.Equal(t, "test", entry.Word)
		assert.Equal(t, "noun", entry.Pos)
		assert.Equal(t, 0, len(entry.Forms))
	})

	t.Run("empty tags", func(t *testing.T) {
		jsonData := `{"form": "test", "tags": []}`
		var wordForm WordForm
		err := json.Unmarshal([]byte(jsonData), &wordForm)
		assert.NoError(t, err)
		assert.Equal(t, "test", wordForm.Form)
		assert.Equal(t, 0, len(wordForm.Tags))
	})
}
