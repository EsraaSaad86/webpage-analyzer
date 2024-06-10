package models

// AnalyzerInfo struct contains information for analysis.
type AnalyzerInfo struct {
	Body string
	Host string
}

// NewAnalyzerInfo creates a new instance of AnalyzerInfo.
func NewAnalyzerInfo(body, host string) AnalyzerInfo {
	return AnalyzerInfo{
		Body: body,
		Host: host,
	}
}
