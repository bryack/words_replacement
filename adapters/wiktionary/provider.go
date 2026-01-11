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

func NewProvider(baseURL string) *Provider {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	wikiClient, err := wiki.NewWikiClient(baseURL, client)
	if err != nil {
		panic(fmt.Sprintf("failed to create wiki client: %v", err))
	}
	return &Provider{
		wikiClient: wikiClient,
	}
}

func (p *Provider) GetForms(word string) (singular, plural []string, err error) {
	return nil, nil, fmt.Errorf("not implemented yet")
}
