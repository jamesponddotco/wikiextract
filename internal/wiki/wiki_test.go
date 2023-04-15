package wiki_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.sr.ht/~jamesponddotco/wikiextract/internal/wiki"
	"github.com/gocolly/colly"
)

func TestNewCrawler(t *testing.T) {
	t.Parallel()

	c := wiki.NewCrawler()
	if c == nil || c.Collector == nil {
		t.Error("NewCrawler did not initialize a Crawler with a non-nil Collector")
	}
}

func TestCrawl(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		handler http.Handler
		wantErr bool
	}{
		{
			name: "Successful crawl",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "Test page")
			}),
			wantErr: false,
		},
		{
			name: "Failed crawl",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(tt.handler)
			defer server.Close()

			c := wiki.NewCrawler()
			err := c.Crawl(server.URL)

			if tt.wantErr && err == nil {
				t.Error("Expected an error, but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
		})
	}
}

func TestOnHTML(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		selector string
		html     string
		expected string
	}{
		{
			name:     "Find element",
			selector: "h1",
			html:     "<html><head></head><body><h1>Test Header</h1></body></html>",
			expected: "Test Header",
		},
		{
			name:     "Element not found",
			selector: "h2",
			html:     "<html><head></head><body><h1>Test Header</h1></body></html>",
			expected: "",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := wiki.NewCrawler()
			var result string
			c.OnHTML(tt.selector, func(e *colly.HTMLElement) {
				result = e.Text
			})

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, tt.html)
			}))
			defer server.Close()

			if err := c.Crawl(server.URL); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected '%v', but got '%v'", tt.expected, result)
			}
		})
	}
}
