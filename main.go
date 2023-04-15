package main

import (
	"os"

	"git.sr.ht/~jamesponddotco/wikiextract/internal/app"
)

func main() {
	os.Exit(app.Run(os.Args[1:]))
}
