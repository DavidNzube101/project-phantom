package domains

import (
	"encoding/json"
	"io/ioutil"
	"time"
	"fmt"
)

// DomainManager manages a collection of domains.
type DomainManager struct {
	Domains   []Domain
	filePath string
}

// NewDomainManager creates a new DomainManager.
func NewDomainManager(filePath string) *DomainManager {
	return &DomainManager{
		filePath: filePath,
	}
}

// LoadDomains loads domains from a JSON file and performs basic validation.
func (dm *DomainManager) LoadDomains() error {
	data, err := ioutil.ReadFile(dm.filePath)
	if err != nil {
		return err
	}

	var rawDomains map[string]map[string]string
	if err := json.Unmarshal(data, &rawDomains); err != nil {
		return err
	}

	dm.Domains = []Domain{}
	for category, sites := range rawDomains {
		for name, url := range sites {
			if !IsValidURL(url) {
				fmt.Printf("Skipping invalid URL: %s for %s\n", url, name)
				continue
			}
			// Basic reachability check - can be made more robust later
			// if !IsReachable(url) {
			// 	fmt.Printf("Skipping unreachable URL: %s for %s\n", url, name)
			// 	continue
			// }

			dm.Domains = append(dm.Domains, Domain{
				URL:            url,
				Name:           name,
				Category:       category,
				LastCrawled:    time.Time{}, // Initialize to zero time
				CrawlFrequency: 48 * time.Hour, // Default to 48 hours
				Priority:       2,             // Default to medium priority
				Status:         "active",
			})
		}
	}

	return nil
}

// GetDomainsByCategory returns a slice of domains belonging to a specific category.
func (dm *DomainManager) GetDomainsByCategory(category string) []Domain {
	var filteredDomains []Domain
	for _, domain := range dm.Domains {
		if domain.Category == category {
			filteredDomains = append(filteredDomains, domain)
		}
	}
	return filteredDomains
}

// GetNextDomainsToCrawl returns domains that need crawling based on their LastCrawled time and CrawlFrequency.
func (dm *DomainManager) GetNextDomainsToCrawl() []Domain {
	var domainsToCrawl []Domain
	now := time.Now()
	for i := range dm.Domains {
		domain := &dm.Domains[i] // Use pointer to modify in place if needed
		if domain.Status == "active" && now.After(domain.LastCrawled.Add(domain.CrawlFrequency)) {
			domainsToCrawl = append(domainsToCrawl, *domain)
		}
	}
	return domainsToCrawl
}

// GetDomainByURL finds a domain by its URL.
func (dm *DomainManager) GetDomainByURL(url string) (*Domain, error) {
	for i := range dm.Domains {
		if dm.Domains[i].URL == url {
			return &dm.Domains[i], nil
		}
	}
	return nil, fmt.Errorf("domain with URL %s not found", url)
}

// UpdateDomain updates an existing domain in the manager.
func (dm *DomainManager) UpdateDomain(updatedDomain *Domain) error {
	for i := range dm.Domains {
		if dm.Domains[i].URL == updatedDomain.URL {
			dm.Domains[i] = *updatedDomain
			return nil
		}
	}
	return fmt.Errorf("domain with URL %s not found for update", updatedDomain.URL)
}

// GetAllDomains returns all the domains.
func (dm *DomainManager) GetAllDomains() []Domain {
	return dm.Domains
}