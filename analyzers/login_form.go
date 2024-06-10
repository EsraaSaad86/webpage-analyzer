package analyzers

import (
	"log"
	"strings"
	"time"
	"webpage-analyzer/analyzers/models"
	"webpage-analyzer/http/responses"

	"golang.org/x/net/html"
)

// LoginFormAnalyzer is an analyzer for detecting login forms.
type LoginFormAnalyzer struct {
}

// NewLoginFormAnalyzer creates a new instance of LoginFormAnalyzer.
func NewLoginFormAnalyzer() Analyzer {
	return &LoginFormAnalyzer{}
}

// Analyze analyzes the web page to detect login forms.
func (a *LoginFormAnalyzer) Analyze(info models.AnalyzerInfo, manager responses.WebPageAnalyzerResponseManager) {
	startTime := time.Now()
	log.Println("LoginFormAnalyzer started")
	defer func(start time.Time) {
		log.Printf("LoginFormAnalyzer completed. Time taken: %v ms", time.Since(start).Milliseconds())
	}(startTime)

	tokenizer := html.NewTokenizer(strings.NewReader(info.Body))
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		} else if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "form" {
				if isLoginForm(tokenizer) {
					manager.SetHasLogin(true)
					return
				}
			}
		}
	}
}

// isLoginForm checks if a <form> element contains input fields commonly found in login forms.
func isLoginForm(tokenizer *html.Tokenizer) bool {
	inputTypes := map[string]bool{"password": true, "text": true, "email": true}

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		} else if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "input" {
				for _, attr := range token.Attr {
					if attr.Key == "type" && inputTypes[attr.Val] {
						return true
					}
				}
			}
		}
	}
	return false
}
