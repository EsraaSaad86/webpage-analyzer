package analyzers

import (
	"log"
	"strings"
	"time"
	"webpage-analyzer/analyzers/models"
	"webpage-analyzer/http/responses"

	"golang.org/x/net/html"
)

// PageTitleAnalyzer is an analyzer for extracting the title of a web page.
type PageTitleAnalyzer struct {
}

// NewPageTitleAnalyzer creates a new instance of PageTitleAnalyzer.
func NewPageTitleAnalyzer() Analyzer {
	return &PageTitleAnalyzer{}
}

// Analyze analyzes the web page to extract the title.
func (a *PageTitleAnalyzer) Analyze(info models.AnalyzerInfo, manager responses.WebPageAnalyzerResponseManager) {
	startTime := time.Now()
	log.Println("PageTitleAnalyzer started")
	defer func(start time.Time) {
		log.Printf("PageTitleAnalyzer completed. Time taken: %v ms", time.Since(start).Milliseconds())
	}(startTime)

	tokenizer := html.NewTokenizer(strings.NewReader(info.Body))
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		} else if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "title" {
				tokenizer.Next()
				title := tokenizer.Token().Data
				manager.SetTitle(title)
				return
			}
		}
	}
}
