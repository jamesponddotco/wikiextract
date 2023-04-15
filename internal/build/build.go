// Package build provides build information about [the wikiextract application].
//
// [the wikiextract application]: https:///git.sr.ht/~jamesponddotco/wikiextract
package build

import "strings"

const (
	// Name is the name of the application.
	Name string = "wikiextract"

	// Version is the version of the application.
	Version string = "0.1.0"

	// URL is the URL of the application.
	URL string = "https:///git.sr.ht/~jamesponddotco/wikiextract"
)

// UserAgent returns the User-Agent string for the application.
func UserAgent() string {
	var builder strings.Builder

	builder.Grow(len(Name) + len(Version) + len(URL) + 20)

	builder.WriteString("Mozilla/5.0 (compatible; ")
	builder.WriteString(Name)
	builder.WriteString("/")
	builder.WriteString(Version)
	builder.WriteString("; +")
	builder.WriteString(URL)
	builder.WriteString(")")

	return builder.String()
}
