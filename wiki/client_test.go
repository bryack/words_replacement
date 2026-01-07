package wiki

import (
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
		if query.Get("action") != "query" {
			t.Fatal("expected action=query")
		}
		if query.Get("titles") != "Go" {
			t.Fatal("expected titles=Go")
		}
		if query.Get("prop") != "extracts" {
			t.Fatal("expected prop=extracts")
		}
		if query.Get("format") != "json" {
			t.Fatal("expected format=json")
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"query":{"pages":{"123":{"extract":"Go is a programming language"}}}}`))
	}))
}

func makeServerWithError(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что запрос корректный
		query := r.URL.Query()
		if query.Get("titles") != "Go" {
			t.Fatal("expected titles=Go")
		}
		w.WriteHeader(http.StatusForbidden)
	}))
}

func makeServerWithEmptyPages(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что запрос корректный
		query := r.URL.Query()
		if query.Get("titles") != "NonExistent" {
			t.Fatal("expected titles=NonExistent")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"query":{"pages":{}}}`))
	}))
}

func makeServerWithSpecialChars(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotTitle := r.URL.Query().Get("titles")
		if gotTitle != "C++" {
			t.Errorf("Server received wrong title: got %q, want 'C++'", gotTitle)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"query":{"pages":{"123":{"extract":"C++ is a programming language"}}}}`))
	}))
}

func TestWikiClient(t *testing.T) {

	t.Run("client can make basic request", func(t *testing.T) {
		svr := makeBasicServer(t)
		defer svr.Close()
		extract, err := GetPage(svr.URL, "Go")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		want := "Go is a programming language"

		if extract != want {
			t.Errorf("want %q, but got %q", want, extract)
		}
	})
	t.Run("handle non 200 status", func(t *testing.T) {
		svr := makeServerWithError(t)
		defer svr.Close()
		extract, err := GetPage(svr.URL, "Go")

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

		extract, err := GetPage(svr.URL, "NonExistent")

		if err == nil {
			t.Fatal("Expected an error when pages map is empty, but got nil")
		}

		if extract != "" {
			t.Errorf("Expected empty extract on error, got %q", extract)
		}

		if !strings.Contains(err.Error(), "page not found") {
			t.Errorf("Expected error message to contain page not found, got %v", err)
		}
	})

	t.Run("Title with special chars", func(t *testing.T) {
		svr := makeServerWithSpecialChars(t)
		defer svr.Close()

		extract, err := GetPage(svr.URL, "C++")

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		want := "C++ is a programming language"

		if extract != want {
			t.Errorf("want %q, but got %q", want, extract)
		}
	})
}
