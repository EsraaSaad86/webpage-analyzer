package responses

type WebPageAnalyzerResponseManager interface {
	SetHtmlVersion(htmlVersion string)
	SetTitle(title string)
	AddHeadingLevel(tag string, level string)
	AddUrlInfo(url string, urlType int, status int, latency int64)
	SetHasLogin(hasLogin bool)
	SetLinkCounts(internalCount, externalCount, inaccessibleCount int)
	ToString() string
	GetAnalysisResponse() AnalysisSuccessResponse
}
