package analyzers

import (
	"webpage-analyzer/analyzers/models"
	"webpage-analyzer/http/responses"
)

// Analyzer defines the interface for webpage analyzers.
type Analyzer interface {
	Analyze(models.AnalyzerInfo, responses.WebPageAnalyzerResponseManager)
}
