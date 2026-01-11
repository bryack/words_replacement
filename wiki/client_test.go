package wiki

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func makeBasicServer(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что запрос корректный
		query := r.URL.Query()
		if query.Get("action") != "parse" {
			t.Fatal("expected action=parse")
		}
		if query.Get("page") != "подделка" {
			t.Fatal("expected page=подделка")
		}
		if query.Get("prop") != "wikitext" {
			t.Fatal("expected prop=wikitext")
		}
		if query.Get("format") != "json" {
			t.Fatal("expected format=json")
		}

		writeJSONResponse(w, `{"parse": {"title": "подделка", "wikitext": {"*": "{{ru-noun|подде́лка|f|подде́лки|подде́лок}}\n\n# [[подделка]]..."}}}`)
	}))
}

func makeServerWithError(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что запрос корректный
		query := r.URL.Query()
		if query.Get("page") != "подделка" {
			t.Fatal("expected page=подделка")
		}
		w.WriteHeader(http.StatusForbidden)
	}))
}

func makeServerWithEmptyPages(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что запрос корректный
		query := r.URL.Query()
		if query.Get("page") != "NonExistent" {
			t.Fatal("expected page=NonExistent")
		}
		w.WriteHeader(http.StatusOK)
		writeJSONResponse(w, `{"parse": {"title": "", "wikitext": {"": ""}}}`)
	}))
}

func makeServerWithSpecialChars(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotTitle := r.URL.Query().Get("page")
		if gotTitle != "C++" {
			t.Errorf("Server received wrong title: got %q, want 'C++'", gotTitle)
		}

		writeJSONResponse(w, `{"parse": {"title": "C++", "wikitext": {"*": "C++ is a programming language"}}}`)
	}))
}

func TestWikiClient(t *testing.T) {

	t.Run("client can make basic request", func(t *testing.T) {
		svr := makeBasicServer(t)
		defer svr.Close()

		wc, err := NewWikiClient(svr.URL, &http.Client{})
		if err != nil {
			t.Fatalf("Failed to create WikiClient: %v", err)
		}
		extract, err := wc.GetPage("подделка")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		want := "{{ru-noun|подде́лка|f|подде́лки|подде́лок}}\n\n# [[подделка]]..."

		if extract != want {
			t.Errorf("want %q, but got %q", want, extract)
		}
	})
	t.Run("handle non 200 status", func(t *testing.T) {
		svr := makeServerWithError(t)
		defer svr.Close()

		wc, err := NewWikiClient(svr.URL, &http.Client{})
		if err != nil {
			t.Fatalf("Failed to create WikiClient: %v", err)
		}
		extract, err := wc.GetPage("подделка")

		if err == nil {
			t.Fatal("Expected an error when server returns 403, but got nil")
		}

		if extract != "" {
			t.Errorf("Expected empty extract on error, got %q", extract)
		}

		if !strings.Contains(err.Error(), "403") {
			t.Errorf("Expected error message to contain status code, got %v", err)
		}
	})
	t.Run("Empty Pages", func(t *testing.T) {
		svr := makeServerWithEmptyPages(t)
		defer svr.Close()

		wc, err := NewWikiClient(svr.URL, &http.Client{})
		if err != nil {
			t.Fatalf("Failed to create WikiClient: %v", err)
		}
		extract, err := wc.GetPage("NonExistent")

		if err == nil {
			t.Fatal("Expected an error when pages map is empty, but got nil")
		}

		if extract != "" {
			t.Errorf("Expected empty extract on error, got %q", extract)
		}

		if !strings.Contains(err.Error(), "no wikitext found for page") {
			t.Errorf("Expected error message to contain page not found, got %v", err)
		}
	})

	t.Run("Title with special chars", func(t *testing.T) {
		svr := makeServerWithSpecialChars(t)
		defer svr.Close()

		wc, err := NewWikiClient(svr.URL, &http.Client{})
		if err != nil {
			t.Fatalf("Failed to create WikiClient: %v", err)
		}
		extract, err := wc.GetPage("C++")

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		want := "C++ is a programming language"

		if extract != want {
			t.Errorf("want %q, but got %q", want, extract)
		}
	})
}

func writeJSONResponse(w http.ResponseWriter, jsonData string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonData))
}

type spyRoundTripper struct {
	called bool
}

func (s *spyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	s.called = true
	if req.URL.Query().Get("page") != "Go" {
		return nil, fmt.Errorf("unexpected title")
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"parse": {"title": "Go", "wikitext": {"*": "{{ru-noun|подде́лка|f|подде́лки|подде́лок}}\n\n# [[подделка]]..."}}}`)),
		Header:     make(http.Header),
	}, nil
}

func TestGetPage_UsesInjectedClient(t *testing.T) {

	t.Run("GetPage should use the provided http.Client", func(t *testing.T) {
		spy := &spyRoundTripper{}
		client := &http.Client{Transport: spy}

		wc, err := NewWikiClient("http://example.com", client)
		if err != nil {
			t.Fatalf("Failed to create WikiClient: %v", err)
		}
		_, err = wc.GetPage("Go")
		if err != nil {
			t.Fatal(err)
		}

		if !spy.called {
			t.Fatal("expected provided http.Client to be used")
		}
	})
}
