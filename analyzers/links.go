package analyzers

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	"webpage-analyzer/analyzers/models"
	"webpage-analyzer/http/responses"

	"golang.org/x/net/html"
)

// LinkType represents the type of a link.
type LinkType int

const (
	Internal LinkType = iota
	External
)

type linkAnalyzer struct{}

// NewLinkAnalyzer creates a new instance of LinkAnalyzer.
func NewLinksAnalyzer() Analyzer {
	return &linkAnalyzer{}
}

// Analyze analyzes the links in a webpage.
func (l *linkAnalyzer) Analyze(data models.AnalyzerInfo, analysis responses.WebPageAnalyzerResponseManager) {
	startTime := time.Now()
	log.Println("Link analyzer started")
	defer func(start time.Time) {
		log.Printf("Link analyzer completed. Time taken: %v ms", time.Since(start).Milliseconds())
	}(startTime)

	links := extractLinks(data.Body, data.Host)

	var wg sync.WaitGroup
	wg.Add(len(links))

	internalCount := 0
	externalCount := 0
	inaccessibleCount := 0

	for _, link := range links {
		go func(link string) {
			defer wg.Done()
			statusCode, latency := checkLink(link)
			analysis.AddUrlInfo(link, int(getLinkType(link, data.Host)), statusCode, latency)

			switch getLinkType(link, data.Host) {
			case Internal:
				internalCount++
			case External:
				externalCount++
			}
			if statusCode < 200 || statusCode >= 400 {
				inaccessibleCount++
			}
		}(link)
	}

	wg.Wait()

	// Set the link counts in the response
	analysis.SetLinkCounts(internalCount, externalCount, inaccessibleCount)
}

// extractLinks extracts links from the HTML content.
func extractLinks(htmlContent, host string) []string {
	var links []string
	tokenizer := html.NewTokenizer(strings.NewReader(htmlContent))

	for {
		switch tokenizer.Next() {
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" || token.Data == "script" || token.Data == "link" {
				var attrName string
				if token.Data == "script" {
					attrName = "src"
				} else if token.Data == "link" {
					attrName = "href"
				} else {
					attrName = "href"
				}
				link := getTagAttribute(token, attrName)
				if isValidLink(link) {
					links = append(links, link)
				}
			}
		case html.ErrorToken:
			return links
		}
	}
}

// getTagAttribute retrieves the value of a specific attribute from a HTML tag.
func getTagAttribute(token html.Token, attrName string) string {
	for _, attr := range token.Attr {
		if attr.Key == attrName {
			return attr.Val
		}
	}
	return ""
}

// checkLink checks the status and latency of a link.
func checkLink(link string) (statusCode int, latency int64) {
	start := time.Now()
	resp, err := http.Get(link)
	if err != nil {
		log.Printf("Error accessing link %s: %v", link, err)
		return http.StatusInternalServerError, 0
	}
	defer resp.Body.Close()
	latency = time.Since(start).Milliseconds()
	return resp.StatusCode, latency
}

// getLinkType identifies the type of the link (internal or external).
func getLinkType(link, host string) LinkType {
	if strings.Contains(link, host) {
		return Internal
	}
	return External
}

// isValidLink checks if the link is valid (starts with "http" or "https").
func isValidLink(link string) bool {
	return strings.HasPrefix(strings.TrimSpace(link), "http://") || strings.HasPrefix(strings.TrimSpace(link), "https://")
}
