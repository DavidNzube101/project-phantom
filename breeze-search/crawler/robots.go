package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/temoto/robotstxt"
)

// RobotsTxtHandler handles robots.txt rules for a given domain.
type RobotsTxtHandler struct {
	robotsData    map[string]*robotstxt.RobotsData
	lastFetched   map[string]time.Time
	fetchInterval time.Duration // How often to re-fetch robots.txt
	mu            sync.Mutex    // Mutex to protect maps
}

// NewRobotsTxtHandler creates a new RobotsTxtHandler.
func NewRobotsTxtHandler() *RobotsTxtHandler {
	return &RobotsTxtHandler{
		robotsData:    make(map[string]*robotstxt.RobotsData),
		lastFetched:   make(map[string]time.Time),
		fetchInterval: 24 * time.Hour, // Default to fetching once a day
	}
}

// CanFetch checks if the given URL can be fetched according to robots.txt rules.
func (r *RobotsTxtHandler) CanFetch(urlStr string) bool {
	u, err := url.Parse(urlStr)
	if err != nil {
		return false // Malformed URL, cannot fetch
	}
	domain := u.Hostname()

	r.mu.Lock()
	data, ok := r.robotsData[domain]
	lastFetched := r.lastFetched[domain]
	r.mu.Unlock()

	// If we don't have robots.txt data or it's stale, try to fetch it.
	if !ok || time.Since(lastFetched) > r.fetchInterval {
		err := r.FetchRobotsTxt(domain)
		if err != nil {
			// If fetching fails, assume we can crawl to be safe, or implement a stricter policy.
			// For now, we'll allow if robots.txt can't be fetched.
			return true
		}
		r.mu.Lock()
		data = r.robotsData[domain]
		r.mu.Unlock()
	}

	if data == nil {
		return true // No robots.txt found or failed to parse, allow crawling
	}

	return data.TestAgent(u.Path, "*") // Test with a generic user agent
}

// FetchRobotsTxt fetches and parses the robots.txt file for a given domain.
func (r *RobotsTxtHandler) FetchRobotsTxt(domain string) error {
	robotsURL := fmt.Sprintf("http://%s/robots.txt", domain)
	resp, err := http.Get(robotsURL)
	if err != nil {
		return fmt.Errorf("failed to fetch robots.txt for %s: %w", domain, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// If robots.txt returns 404 or other non-200 status, it means no robots.txt exists.
		// In this case, we treat it as allowing all.
		if resp.StatusCode == http.StatusNotFound {
			r.mu.Lock()
			r.robotsData[domain] = nil // Mark as no robots.txt
			r.lastFetched[domain] = time.Now()
			r.mu.Unlock()
			return nil
		}
		return fmt.Errorf("received non-OK status code %d for robots.txt from %s", resp.StatusCode, domain)
	}

	data, err := robotstxt.FromResponse(resp)
	if err != nil {
		return fmt.Errorf("failed to parse robots.txt for %s: %w", domain, err)
	}

	r.mu.Lock()
	r.robotsData[domain] = data
	r.lastFetched[domain] = time.Now()
	r.mu.Unlock()

	return nil
}