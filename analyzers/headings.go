package analyzers

import (
	"log"
	"regexp"
	"strings"
	"time"
	"webpage-analyzer/analyzers/models"
	"webpage-analyzer/http/responses"
)

// HeaderAnalyzer is an analyzer for detecting header tags.
type HeaderAnalyzer struct {
}

// NewHeaderAnalyzer creates a new instance of HeaderAnalyzer.
func NewHeaderAnalyzer() Analyzer {
	return &HeaderAnalyzer{}
}

// Analyze analyzes the web page to detect header tags.
func (h *HeaderAnalyzer) Analyze(info models.AnalyzerInfo, manager responses.WebPageAnalyzerResponseManager) {
	startTime := time.Now()
	log.Println("HeaderAnalyzer started")
	defer func(start time.Time) {
		log.Printf("HeaderAnalyzer completed. Time taken: %v ms", time.Since(start).Milliseconds())
	}(startTime)

	headerRegex := regexp.MustCompile(`(?i)<[hH][1-9][^>]*>.*?<\/[hH][1-9]>`)

	// Find all matches of header tags in the HTML content
	headerMatches := headerRegex.FindAllString(info.Body, -1)

	// Iterate over each header match and extract the header text
	for _, match := range headerMatches {
		headerTag := extractHeaderTag(match)   // Extract header tag (e.g., <h1>, <h2>, etc.)
		headerText := extractHeaderText(match) // Extract header text between <h1> and </h1>

		// Add header information to the response manager
		manager.AddHeadingLevel(headerTag, headerText)
	}
}

// extractHeaderTag extracts the header tag (e.g., <h1>, <h2>, etc.) from the header match.
func extractHeaderTag(headerMatch string) string {
	tagStart := strings.IndexAny(headerMatch, "hH")                 // Find the start index of the header tag
	tagEnd := strings.Index(headerMatch[tagStart:], ">") + tagStart // Find the end index of the header tag
	if tagStart != -1 && tagEnd != -1 {
		return headerMatch[tagStart : tagEnd+1] // Extract the header tag substring
	}
	return ""
}

// extractHeaderText extracts the text between <h1> and </h1> tags.
func extractHeaderText(headerMatch string) string {
	startIndex := strings.Index(headerMatch, ">") + 1 // Find the start index of the header text
	endIndex := strings.LastIndex(headerMatch, "<")   // Find the end index of the header text
	if startIndex >= 0 && endIndex >= 0 && startIndex < endIndex {
		return headerMatch[startIndex:endIndex]
	}
	return ""
}
