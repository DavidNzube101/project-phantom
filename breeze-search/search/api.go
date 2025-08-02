package search

import (
	"encoding/json"
	"net/http"
)

// API represents the REST API server for the search engine.
type API struct {
	Searcher *Searcher
}

// NewAPI creates a new API instance.
func NewAPI(searcher *Searcher) *API {
	return &API{
		Searcher: searcher,
	}
}

// RegisterRoutes registers the API endpoints.
func (api *API) RegisterRoutes() {
	http.HandleFunc("/search", api.handleSearch)
}

// handleSearch handles search requests.
func (api *API) handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	results := api.Searcher.Search(query)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
