// Package wiki provides functions to crawl and extract information from
// MediaWiki websites.
package wiki

import (
	"fmt"

	"git.sr.ht/~jamesponddotco/wikiextract/internal/build"
	"github.com/gocolly/colly"
)

// Crawler is a MediaWiki crawler. This is mostly a wrapper around [Colly].
//
// [Colly]: https://github.com/gocolly/colly
type Crawler struct {
	// Collector is the Colly collector used to crawl the website.
	Collector *colly.Collector
}

func NewCrawler() *Crawler {
	return &Crawler{
		Collector: colly.NewCollector(
			colly.MaxDepth(1),
			colly.UserAgent(build.UserAgent()),
		),
	}
}

func (c *Crawler) OnHTML(selector string, fun colly.HTMLCallback) {
	c.Collector.OnHTML(selector, fun)
}

func (c *Crawler) Crawl(url string) error {
	if err := c.Collector.Visit(url); err != nil {
		return fmt.Errorf("error visiting %s: %w", url, err)
	}

	return nil
}
