package sqlite

// KaikkiEntry represents a single entry from Kaikki.org JSONL dictionary.
type KaikkiEntry struct {
	Word  string     `json:"word"`
	Pos   string     `json:"pos"`
	Forms []WordForm `json:"forms"`
}

// WordForms represents a single word form with grammatical tags.
type WordForm struct {
	Form string   `json:"form"`
	Tags []string `json:"tags"`
}
