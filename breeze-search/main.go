package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"breeze-search/core"
	"breeze-search/crawler"
	"breeze-search/domains"
	"breeze-search/indexer"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 1. Load domains using your DomainManager
	dm := domains.NewDomainManager("domains/solana_sites.json")
	err := dm.LoadDomains()
	if err != nil {
		log.Fatalf("Failed to load domains: %v", err)
	}

	loadedDomains := dm.GetAllDomains()
	if len(loadedDomains) == 0 {
		log.Fatal("No domains loaded. Please ensure domains/solana_sites.json has content.")
	}

	log.Printf("Loaded %d domains.", len(loadedDomains))

	// 2. Create the Indexer
	idx := indexer.NewIndexer()

	// 3. Configure and start the Crawler
	crawlerConfig := &crawler.CrawlerConfig{
		MaxDepth:           1, // Only crawl the initial page for now
		MaxConcurrentPages: 5, // Number of concurrent workers
		RequestTimeout:     30 * time.Second,
		UserAgent:          "BreezeSearchBot/1.0",
	}

	c := crawler.NewCrawler(crawlerConfig, dm, idx)
	c.Start()

	// Enqueue the first 5-10 domains for crawling
	numToCrawl := 5
	if len(loadedDomains) < numToCrawl {
		numToCrawl = len(loadedDomains)
	}

	log.Printf("Enqueuing %d domains for crawling...", numToCrawl)
	for i := 0; i < numToCrawl; i++ {
		domain := loadedDomains[i]
		// For simplicity, we'll just crawl the base URL of the domain.
		// In a real scenario, you might want to construct a proper starting URL.
		job := core.CrawlJob{
			URL:    fmt.Sprintf("http://%s", domain.Name),
			Domain: domain.Name,
		}
		c.EnqueueJob(job)
	}

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan // Wait for interrupt signal

	log.Println("Shutting down crawler...")
	c.Stop()
	log.Println("Crawler shut down.")

	// Print the number of indexed documents
	log.Printf("Total documents indexed: %d", idx.GetDocumentCount())

	log.Println("Exiting.")
}
