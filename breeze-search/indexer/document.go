package indexer

import (
	"time"
)

// Document represents a single indexed document.
type Document struct {
	ID          string    `json:"id"` // Unique ID, e.g., URL hash
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	BodyText    string    `json:"body_text"`
	Links       []string  `json:"links"`
	Timestamp   time.Time `json:"timestamp"`
	ContentHash string    `json:"content_hash"`
}
