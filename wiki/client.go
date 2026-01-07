package wiki

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetPage(baseURL, title string) (string, error) {
	requestURL, err := buildWikiURL(baseURL, title)
	if err != nil {
		return "", err
	}

	body, err := makeWikiRequest(requestURL)
	if err != nil {
		return "", err
	}

	return parseWikiResponse(body)
}

func buildWikiURL(baseURL, title string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}
	q := u.Query()
	q.Set("action", "query")
	q.Set("titles", title)
	q.Set("prop", "extracts")
	q.Set("format", "json")
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func makeWikiRequest(requestURL string) ([]byte, error) {
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func parseWikiResponse(body []byte) (string, error) {
	var result struct {
		Query struct {
			Pages map[string]struct {
				Extract string `json:"extract"`
			} `json:"pages"`
		} `json:"query"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	for _, page := range result.Query.Pages {
		return page.Extract, nil
	}
	return "", fmt.Errorf("error: page not found")
}
