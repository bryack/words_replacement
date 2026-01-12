package wiki

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type WikiClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewWikiClient(baseURL string, httpClient *http.Client) (*WikiClient, error) {
	if baseURL == "" {
		return nil, fmt.Errorf("baseURL must not be empty")
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 5 * time.Second,
		}
	}
	return &WikiClient{baseURL: baseURL, httpClient: httpClient}, nil
}

func (wc *WikiClient) GetPage(title string) (string, error) {
	requestURL, err := wc.buildWikiURL(title)
	if err != nil {
		return "", fmt.Errorf("failed to build wiki URL: %w", err)
	}

	body, err := wc.makeWikiRequest(requestURL)
	if err != nil {
		return "", fmt.Errorf("error: not found: %w", err)
	}

	return parseWikiResponse(body)
}

func (wc *WikiClient) buildWikiURL(title string) (string, error) {
	u, err := url.Parse(wc.baseURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL %q: %w", wc.baseURL, err)
	}
	q := u.Query()
	q.Set("action", "parse")
	q.Set("page", title)
	q.Set("prop", "wikitext")
	q.Set("format", "json")
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (wc *WikiClient) makeWikiRequest(requestURL string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "WordsBot/1.0 (https://github.com/bryack/words)")
	resp, err := wc.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request %q: %w", requestURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request %q failed with status: %d", requestURL, resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func parseWikiResponse(body []byte) (string, error) {
	var result struct {
		Parse struct {
			Title    string `json:"title"`
			Wikitext struct {
				Content string `json:"*"`
			} `json:"wikitext"`
		} `json:"parse"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to unmarshal wiki response: %w", err)
	}

	if result.Parse.Wikitext.Content == "" {
		return "", fmt.Errorf("no wikitext found for page")
	}
	return result.Parse.Wikitext.Content, nil
}
