package wiktionary

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bryack/words/wiki"
)

const defaultTimeout = 10 * time.Second

type Provider struct {
	wikiClient *wiki.WikiClient
	parser     *RussianNounParser
}

func NewProvider(baseURL string) (*Provider, error) {
	client := &http.Client{
		Timeout: defaultTimeout,
	}
	wikiClient, err := wiki.NewWikiClient(baseURL, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create wiki client: %w", err)
	}
	return &Provider{
		wikiClient: wikiClient,
		parser:     NewRussianNounParser(),
	}, nil
}

func (p *Provider) GetForms(word string) (singular, plural []string, err error) {
	wikitext, err := p.wikiClient.GetPage(word)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get wikitext for %q: %w", word, err)
	}

	return p.parser.ParseForms(wikitext)
}
