package indexer

import (
	"sync"

	"breeze-search/core"
)

// Indexer manages the storage and retrieval of documents.
type Indexer struct {
	// For now, a simple in-memory map. This will be replaced by BadgerDB/BoltDB.
	Documents map[string]Document
	Mu        sync.RWMutex // Mutex for concurrent access to documents map
}

// NewIndexer creates a new Indexer instance.
func NewIndexer() *Indexer {
	return &Indexer{
		Documents: make(map[string]Document),
	}
}

// AddDocument adds a new document to the index.
func (idx *Indexer) AddDocument(result core.CrawlResult) {
	idx.Mu.Lock()
	defer idx.Mu.Unlock()

	doc := Document{
		ID:          result.URL, // Using URL as ID for simplicity for now
		URL:         result.URL,
		Title:       result.Title,
		Description: result.Description,
		BodyText:    result.BodyText,
		Links:       result.Links,
		Timestamp:   result.Timestamp,
		ContentHash: result.ContentHash,
	}

	idx.Documents[doc.ID] = doc
}

// GetDocument retrieves a document by its ID.
func (idx *Indexer) GetDocument(id string) (Document, bool) {
	idx.Mu.RLock()
	defer idx.Mu.RUnlock()

	doc, ok := idx.Documents[id]
	return doc, ok
}

// GetDocumentCount returns the number of documents in the index.
func (idx *Indexer) GetDocumentCount() int {
	idx.Mu.RLock()
	defer idx.Mu.RUnlock()
	return len(idx.Documents)
}
