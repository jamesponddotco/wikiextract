// Package extract provides functions to extract words from a string.
package extract

import (
	"regexp"
	"strings"
)

// Words extracts and appends words from a string to a string slice given a
// regular expression.
func Words(str string, re *regexp.Regexp, slice []string) []string {
	words := re.FindAllString(str, -1)

	for _, word := range words {
		lowercaseWord := strings.ToLower(word)

		if !contains(slice, lowercaseWord) {
			slice = append(slice, lowercaseWord)
		}
	}

	return slice
}

// contains checks if a string slice contains a string.
func contains(slice []string, word string) bool {
	for _, s := range slice {
		if s == word {
			return true
		}
	}

	return false
}
