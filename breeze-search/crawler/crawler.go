package crawler

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"breeze-search/core"
	"breeze-search/domains"
	"breeze-search/indexer"
)

// CrawlerConfig holds configuration settings for the crawler.
type CrawlerConfig struct {
	MaxDepth           int
	MaxConcurrentPages int
	RequestTimeout     time.Duration
	UserAgent          string
	// Add more configuration as needed, e.g., rate limits per domain
}

// Crawler orchestrates the crawling process.
type Crawler struct {
	config        *CrawlerConfig
	domainManager *domains.DomainManager
	indexer       *indexer.Indexer
	jobQueue      chan core.CrawlJob
	resultChan    chan core.CrawlResult
	workers       []*Worker
	quit          chan bool
	wg            sync.WaitGroup
	extractor     *Extractor
	robotsHandler *RobotsTxtHandler
}

// NewCrawler creates a new Crawler instance.
func NewCrawler(config *CrawlerConfig, dm *domains.DomainManager, idx *indexer.Indexer) *Crawler {
	return &Crawler{
		config:        config,
		domainManager: dm,
		indexer:       idx,
		jobQueue:      make(chan core.CrawlJob, config.MaxConcurrentPages),
		resultChan:    make(chan core.CrawlResult, config.MaxConcurrentPages),
		quit:          make(chan bool),
		extractor:     NewExtractor(),
		robotsHandler: NewRobotsTxtHandler(),
	}
}

// Start begins the crawling process.
func (c *Crawler) Start() {
	log.Println("Crawler starting...")

	// Start workers
	for i := 0; i < c.config.MaxConcurrentPages; i++ {
		worker := NewWorker(i+1, c.jobQueue, c.resultChan, c)
		c.workers = append(c.workers, worker)
		worker.Start()
	}

	// Goroutine to process crawl results
	go c.processResults()

	log.Printf("Crawler started with %d workers.", len(c.workers))
}

// Stop gracefully stops the crawling process.
func (c *Crawler) Stop() {
	log.Println("Crawler stopping...")
	close(c.quit)

	// Signal workers to stop
	for _, worker := range c.workers {
		worker.Stop()
	}

	// Wait for all active crawls to finish
	c.wg.Wait()
	close(c.jobQueue)
	close(c.resultChan)
	log.Println("Crawler stopped.")
}

// EnqueueJob adds a new crawl job to the queue.
func (c *Crawler) EnqueueJob(job core.CrawlJob) {
	c.wg.Add(1)
	c.jobQueue <- job
}

// processResults handles the results coming from workers.
func (c *Crawler) processResults() {
	for {
		select {
		case result := <-c.resultChan:
			c.handleCrawlResult(result)
			c.wg.Done()
		case <-c.quit:
			log.Println("Result processor stopping.")
			return
		}
	}
}

// handleCrawlResult processes a single crawl result.
func (c *Crawler) handleCrawlResult(result core.CrawlResult) {
	if result.Success {
		log.Printf("Successfully crawled: %s", result.URL)
		// Update LastCrawled timestamp in DomainManager
		domain, err := c.domainManager.GetDomainByURL(result.URL)
		if err != nil {
			log.Printf("Error getting domain for URL %s: %v", result.URL, err)
			return
		}
		domain.LastCrawled = result.LastCrawled
		err = c.domainManager.UpdateDomain(domain)
		if err != nil {
			log.Printf("Error updating domain %s: %v", domain.Name, err)
		}

		// Pass extracted content to indexer
		c.indexer.AddDocument(result)
		log.Printf("Indexed document from: %s", result.URL)

	} else {
		log.Printf("Failed to crawl %s: %v", result.URL, result.Error)
		// Handle crawl failure gracefully, e.g., update domain status
	}
}

// CrawlURL performs the actual HTTP request and content extraction.
func (c *Crawler) CrawlURL(job core.CrawlJob) core.CrawlResult {
	result := core.CrawlResult{
		URL: job.URL,
		Success: false,
		LastCrawled: time.Now(),
	}

	// Check robots.txt
	if !c.robotsHandler.CanFetch(job.URL) {
		result.Error = fmt.Errorf("crawling disallowed by robots.txt")
		return result
	}

												collyCollector := colly.NewCollector(
		colly.MaxDepth(c.config.MaxDepth),
		colly.UserAgent(c.config.UserAgent),
		colly.AllowURLRevisit(), // Allow crawling the same URL multiple times
	)

	// Set request timeout
	collyCollector.SetClient(&http.Client{
		Timeout: c.config.RequestTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // For development, ignore invalid certs
		},
	})

	// Randomize user agent (optional, but good practice)
	extensions.RandomUserAgent(collyCollector)

	// Set error handler
	collyCollector.OnError(func(r *colly.Response, err error) {
		result.Error = fmt.Errorf("request failed: %w", err)
		result.StatusCode = r.StatusCode
		log.Printf("Error crawling %s: %v (Status: %d)", r.Request.URL.String(), err, r.StatusCode)
	})

	// Set HTML callback for content extraction
	collyCollector.OnHTML("html", func(e *colly.HTMLElement) {
		htmlContent, err := e.DOM.Html()
		if err != nil {
			result.Error = fmt.Errorf("failed to get HTML content: %w", err)
			return
		}

				title, description, bodyText, links, contentHash, err := c.extractor.ExtractContent([]byte(htmlContent))
		if err != nil {
			result.Error = fmt.Errorf("failed to extract content: %w", err)
			return
		}

		result.Title = title
		result.Description = description
		result.BodyText = bodyText
		result.Links = links
		result.ContentHash = contentHash
		result.Success = true
		result.StatusCode = e.Response.StatusCode
	})

	// Visit the URL
	err := collyCollector.Visit(job.URL)
	if err != nil {
		// If error was already set by OnError, don't overwrite
		if result.Error == nil {
			result.Error = fmt.Errorf("colly visit error: %w", err)
		}
	}

	collyCollector.Wait() // Wait for the crawl to finish

	return result
}


