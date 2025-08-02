package crawler

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Extractor is responsible for extracting and cleaning content from HTML.
type Extractor struct{}

// NewExtractor creates a new Extractor.
func NewExtractor() *Extractor {
	return &Extractor{}
}

// ExtractContent extracts title, description, body text, and links from HTML.
func (e *Extractor) ExtractContent(htmlContent []byte) (title, description, bodyText string, links []string, contentHash string, err error) {
	reader := bytes.NewReader(htmlContent)
	doc, err := html.Parse(reader)
	if err != nil {
		return "", "", "", nil, "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	title = e.extractTitle(doc)
	description = e.extractMetaDescription(doc)
	bodyText = e.extractBodyText(doc) // Pass doc to extractBodyText
	links = e.extractLinks(doc)
	contentHash = e.generateContentHash(htmlContent)

	return title, description, bodyText, links, contentHash, nil
}

func (e *Extractor) extractTitle(n *html.Node) string {
	var title string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
			if title != "" {
				return
			}
		}
	}
	f(n)
	return title
}

func (e *Extractor) extractMetaDescription(n *html.Node) string {
	var description string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			var name, content string
			for _, attr := range n.Attr {
				if attr.Key == "name" && attr.Val == "description" {
					name = attr.Val
				} else if attr.Key == "content" {
					content = attr.Val
				}
			}
			if name == "description" && content != "" {
				description = content
				return
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
			if description != "" {
				return
			}
		}
	}
	f(n)
	return description
}

func (e *Extractor) extractBodyText(doc *html.Node) string {
	var buf bytes.Buffer
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			// Clean up text: remove excessive whitespace, newlines
			text := strings.TrimSpace(n.Data)
			if text != "" {
				buf.WriteString(text)
				buf.WriteString(" ") // Add space between text nodes
			}
		}
		// Skip script and style elements
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style" || n.Data == "nav" || n.Data == "footer" || n.Data == "header" || n.Data == "aside") {
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	// Start extraction from the body tag to avoid extracting head content
	var bodyNode *html.Node
	var findBody func(*html.Node)
	findBody = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			bodyNode = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findBody(c)
			if bodyNode != nil {
				return
			}
		}
	}
	findBody(doc)

	if bodyNode != nil {
		f(bodyNode)
	}
	return strings.Join(strings.Fields(buf.String()), " ") // Normalize whitespace
}

func (e *Extractor) extractLinks(n *html.Node) []string {
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					// TODO: Resolve relative URLs to absolute URLs
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return links
}

func (e *Extractor) generateContentHash(content []byte) string {
	hash := sha256.Sum256(content)
	return fmt.Sprintf("%x", hash)
}
