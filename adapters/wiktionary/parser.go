package wiktionary

import (
	"fmt"
	"regexp"
	"strings"
)

var ruNounRegex = regexp.MustCompile(`\{\{ru-noun\|([^|]*)\|[^|]*\|([^|]*)\|([^|}]*)`)

const minMatchGroups = 4

type RussianNounParser struct{}

func NewRussianNounParser() *RussianNounParser {
	return &RussianNounParser{}
}

func (r *RussianNounParser) ParseForms(wikitext string) (singular, plural []string, err error) {
	matches := ruNounRegex.FindStringSubmatch(wikitext)

	if len(matches) < minMatchGroups {
		return nil, nil, fmt.Errorf("not found")
	}

	for _, v := range matches[1:] {
		if v == "" {
			return nil, nil, fmt.Errorf("empty parameter")
		}
	}

	singular = append(singular, strings.TrimSpace(matches[1]))
	plural = append(plural, strings.TrimSpace(matches[2]), strings.TrimSpace(matches[3]))
	return singular, plural, nil
}
