package search

import (
	"strings"

	"breeze-search/indexer"
)

// Searcher provides search capabilities over indexed documents.
type Searcher struct {
	Indexer *indexer.Indexer
}

// NewSearcher creates a new Searcher instance.
func NewSearcher(idx *indexer.Indexer) *Searcher {
	return &Searcher{
		Indexer: idx,
	}
}

// Search performs a simple keyword search across indexed documents.
// It currently searches in Title, Description, and BodyText.
func (s *Searcher) Search(query string) []indexer.Document {
	var results []indexer.Document
	query = strings.ToLower(query)

	s.Indexer.Mu.RLock() // Acquire read lock for accessing documents
	defer s.Indexer.Mu.RUnlock()

	for _, doc := range s.Indexer.Documents {
		if strings.Contains(strings.ToLower(doc.Title), query) ||
			strings.Contains(strings.ToLower(doc.Description), query) ||
			strings.Contains(strings.ToLower(doc.BodyText), query) {
			results = append(results, doc)
		}
	}
	return results
}
