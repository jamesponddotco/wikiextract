// Package app is where the core logic for wikiextract lives.
package app

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"git.sr.ht/~jamesponddotco/wikiextract/internal/extract"
	"git.sr.ht/~jamesponddotco/wikiextract/internal/wiki"
	"github.com/gocolly/colly"
)

func Run(args []string) int {
	var (
		url      string
		file     string
		listFile string
	)

	flags := flag.NewFlagSet("wikiextract", flag.ExitOnError)
	flags.StringVar(&url, "url", "", "URL of the Wikipedia page")
	flags.StringVar(&url, "u", "", "URL of the Wikipedia page (shorthand)")
	flags.StringVar(&file, "output-file", "", "Output file path")
	flags.StringVar(&file, "o", "", "Output file path (shorthand)")
	flags.StringVar(&listFile, "list-file", "", "File containing a list of URLs")
	flags.StringVar(&listFile, "l", "", "File containing a list of URLs (shorthand)")

	if err := flags.Parse(args); err != nil {
		log.Fatal(err)
		return 1
	}

	if (url == "" && listFile == "") || file == "" {
		log.Fatal("URL or list-file, and output-file must be provided.")
		return 1
	}

	if url != "" && listFile != "" {
		log.Fatal("URL and list-file cannot be used together.")
		return 1
	}

	var urls []string

	if url != "" {
		urls = append(urls, url)
	} else {
		// Read URLs from the list file
		urlsFromFile, err := readURLsFromFile(listFile)
		if err != nil {
			log.Fatal(err)
			return 1
		}
		urls = urlsFromFile
	}

	wordSlice := []string{}
	re := regexp.MustCompile(`(?i)\b[[:alpha:]]{4,}\b`)

	for _, url := range urls {
		crawler := wiki.NewCrawler()

		crawler.OnHTML("div#mw-content-text", func(e *colly.HTMLElement) {
			e.ForEach("p", func(_ int, e *colly.HTMLElement) {
				wordSlice = extract.Words(e.Text, re, wordSlice)
			})
		})

		if err := crawler.Crawl(url); err != nil {
			log.Fatal(err)
			return 1
		}
	}

	if err := writeWordsToFile(wordSlice, file); err != nil {
		log.Fatal(err)
		return 1
	}

	return 0
}

func writeWordsToFile(words []string, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	for _, word := range words {
		if _, err := fmt.Fprintln(f, word); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}

	return nil
}

func readURLsFromFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	urls := []string{}

	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %w", err)
	}

	return urls, nil
}
