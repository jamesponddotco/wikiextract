package build_test

import (
	"strings"
	"testing"

	"git.sr.ht/~jamesponddotco/wikiextract/internal/build"
)

func TestUserAgent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "UserAgent format",
			expected: "Mozilla/5.0 (compatible; " + build.Name + "/" + build.Version + "; +" + build.URL + ")",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ua := build.UserAgent()

			if tt.expected != ua {
				t.Errorf("Expected: %s, got: %s", tt.expected, ua)
			}

			if !strings.Contains(ua, build.Name) || !strings.Contains(ua, build.Version) || !strings.Contains(ua, build.URL) {
				t.Error("UserAgent does not contain required fields")
			}
		})
	}
}
