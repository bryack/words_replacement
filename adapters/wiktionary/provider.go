package wiktionary

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bryack/words/wiki"
)

type Provider struct {
	wikiClient *wiki.WikiClient
}

func NewProvider(baseURL string) (*Provider, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	wikiClient, err := wiki.NewWikiClient(baseURL, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create wiki client: %w", err)
	}
	return &Provider{
		wikiClient: wikiClient,
	}, nil
}

func (p *Provider) GetForms(word string) (singular, plural []string, err error) {
	return nil, nil, fmt.Errorf("not implemented yet")
}
