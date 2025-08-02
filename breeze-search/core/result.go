package core

import (
	"time"
)

// CrawlResult holds the extracted content and metadata from a crawled page.
type CrawlResult struct {
	URL         string
	Success     bool
	StatusCode  int
	Error       error
	Title       string
	Description string
	BodyText    string
	Links       []string
	ContentHash string
	LastCrawled time.Time
	Timestamp   time.Time
}

// CrawlJob represents an individual crawl task.
type CrawlJob struct {
	URL    string
	Domain string
}
