package extract_test

import (
	"regexp"
	"testing"

	"git.sr.ht/~jamesponddotco/wikiextract/internal/extract"
)

func TestWords(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		re     *regexp.Regexp
		slice  []string
		expect []string
	}{
		{
			name:   "No words",
			input:  "",
			re:     regexp.MustCompile(`(?i)\b[[:alpha:]]{4,}\b`),
			slice:  []string{},
			expect: []string{},
		},
		{
			name:   "One word",
			input:  "Wikipedia",
			re:     regexp.MustCompile(`(?i)\b[[:alpha:]]{4,}\b`),
			slice:  []string{},
			expect: []string{"wikipedia"},
		},
		{
			name:   "Multiple words",
			input:  "The quick brown fox jumps over the lazy dog",
			re:     regexp.MustCompile(`(?i)\b[[:alpha:]]{4,}\b`),
			slice:  []string{},
			expect: []string{"quick", "brown", "jumps", "over", "lazy"},
		},
		{
			name:   "Duplicate words",
			input:  "The quick brown dog jumps over the lazy brown dog",
			re:     regexp.MustCompile(`(?i)\b[[:alpha:]]{4,}\b`),
			slice:  []string{},
			expect: []string{"quick", "brown", "jumps", "over", "lazy"},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := extract.Words(tt.input, tt.re, tt.slice)
			if len(result) != len(tt.expect) {
				t.Fatalf("Expected %v, but got %v", tt.expect, result)
			}

			for i, word := range result {
				if word != tt.expect[i] {
					t.Errorf("Expected word %v, but got %v", tt.expect[i], word)
				}
			}
		})
	}
}
