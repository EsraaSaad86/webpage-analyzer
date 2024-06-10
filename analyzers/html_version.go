package analyzers

import (
	"strings"
	"webpage-analyzer/analyzers/models"
	"webpage-analyzer/http/responses"
)

// HTMLVersionAnalyzer struct
type HTMLVersionAnalyzer struct {
	types map[string]string
}

// NewHTMLVersionAnalyzer creates a new instance of HTMLVersionAnalyzer.
func NewHTMLVersionAnalyzer() Analyzer {
	return &HTMLVersionAnalyzer{
		types: map[string]string{
			"HTML 2.0":                `"-//IETF//DTD HTML 2.0//EN"`,
			"HTML 3.2":                `"-//W3C//DTD HTML 3.2 Final//EN"`,
			"HTML 4.01 Strict":        `"-//W3C//DTD HTML 4.01//EN"`,
			"HTML 4.01 Transitional":  `"-//W3C//DTD HTML 4.01 Transitional//EN"`,
			"HTML 4.01 Frameset":      `"-//W3C//DTD HTML 4.01 Frameset//EN"`,
			"XHTML 1.0 Strict":        `"-//W3C//DTD XHTML 1.0 Strict//EN"`,
			"XHTML 1.0 Transitional":  `"-//W3C//DTD XHTML 1.0 Transitional//EN"`,
			"XHTML 1.0 Frameset":      `"-//W3C//DTD XHTML 1.0 Frameset//EN"`,
			"XHTML 1.1":               `"-//W3C//DTD XHTML 1.1//EN"`,
			"HTML 5":                  `<!DOCTYPE html>`,
			"HTML 5.1":                `<!DOCTYPE html>`,
			"HTML 5.2":                `<!DOCTYPE html>`,
			"HTML 5.3":                `<!DOCTYPE html>`,
			"Custom Document Type #1": `<!-- Custom Doctype #1 -->`,
			"Custom Document Type #2": `<!-- Custom Doctype #2 -->`,
			// Add more HTML versions or custom document types as needed
		},
	}
}

// Analyze method to extract HTML version from the web page content.
func (h *HTMLVersionAnalyzer) Analyze(info models.AnalyzerInfo, resManager responses.WebPageAnalyzerResponseManager) {
	body := info.Body
	for version, doctype := range h.types {
		if strings.Contains(body, doctype) {
			resManager.SetHtmlVersion(version)
			return
		}
	}
	resManager.SetHtmlVersion("Unknown")
}
