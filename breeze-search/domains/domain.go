package domains

import (
	"time"
)

// Domain represents a single Solana ecosystem site to be crawled.
type Domain struct {
	URL           string        `json:"url"`
	Name          string        `json:"name"`
	Category      string        `json:"category"`
	LastCrawled   time.Time     `json:"last_crawled"`
	CrawlFrequency time.Duration `json:"crawl_frequency"` // e.g., 24h, 48h, 7d
	Priority      int           `json:"priority"`        // e.g., 1 (high), 2 (medium), 3 (low)
	Status        string        `json:"status"`          // e.g., "active", "failed", "disabled"
}
