package sqlite

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const initialCapacity = 10

// LoadFromJSONL reads JSONL file line by line, parses JSON, filters for nouns only
func LoadFromJSONL(filepath string) ([]KaikkiEntry, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %q: %w", filepath, err)
	}
	defer file.Close()
	entries := make([]KaikkiEntry, 0, initialCapacity)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("empty line")
			continue
		}
		var entry KaikkiEntry
		if err = json.Unmarshal([]byte(line), &entry); err != nil {
			fmt.Printf("failed to unmarshal line: %v\n", err)
			continue
		}
		if entry.Pos == "noun" {
			entries = append(entries, entry)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file %q: %w", filepath, err)
	}
	return entries, nil
}
