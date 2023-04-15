package extract

import "testing"

func TestContains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		slice  []string
		word   string
		expect bool
	}{
		{
			name:   "Empty slice",
			slice:  []string{},
			word:   "example",
			expect: false,
		},
		{
			name:   "Word not in slice",
			slice:  []string{"one", "two", "three"},
			word:   "four",
			expect: false,
		},
		{
			name:   "Word in slice",
			slice:  []string{"one", "two", "three"},
			word:   "two",
			expect: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := contains(tt.slice, tt.word)
			if result != tt.expect {
				t.Errorf("Expected %v, but got %v", tt.expect, result)
			}
		})
	}
}
