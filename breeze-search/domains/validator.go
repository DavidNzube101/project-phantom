package domains

import (
	"net/http"
	"net/url"
	"time"
)

// IsValidURL checks if a string is a valid URL.
func IsValidURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}

// IsReachable checks if a URL is reachable via a HEAD request.
// It follows redirects and returns true if the final status code is 2xx.
func IsReachable(urlStr string) bool {
	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Allow up to 10 redirects
			if len(via) >= 10 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}

	req, err := http.NewRequest("HEAD", urlStr, nil)
	if err != nil {
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode >= 200 && resp.StatusCode < 300
}
