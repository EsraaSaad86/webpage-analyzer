package responses

import (
	"encoding/json"
	"log"
	"sync"
)

type AnalysisSuccessResponseManager struct {
	successRes AnalysisSuccessResponse
	lock       sync.RWMutex
}

type AnalysisSuccessResponse struct {
	HtmlVersion           string    `json:"htmlVersion"`
	Title                 string    `json:"title"`
	Headings              []Heading `json:"headings"`
	Urls                  []Url     `json:"urls"`
	HasLogin              bool      `json:"hasLogin"`
	InternalLinkCount     int       `json:"internalLinkCount"`
	ExternalLinkCount     int       `json:"externalLinkCount"`
	InaccessibleLinkCount int       `json:"inaccessibleLinkCount"`
}

type Heading struct {
	TagName string   `json:"tagName"`
	Levels  []string `json:"levels"`
}

type Url struct {
	Url     string `json:"url"`
	Type    string `json:"type"`
	Status  int    `json:"status"`
	Latency int64  `json:"latency"`
}

func NewAnalysisSuccessResponseManager() *AnalysisSuccessResponseManager {
	return &AnalysisSuccessResponseManager{
		successRes: AnalysisSuccessResponse{},
	}
}

func (w *AnalysisSuccessResponseManager) SetHtmlVersion(htmlVersion string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.HtmlVersion = htmlVersion
}

func (w *AnalysisSuccessResponseManager) SetTitle(title string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.Title = title
}

func (w *AnalysisSuccessResponseManager) AddHeadingLevel(tag string, level string) {
	w.lock.Lock()
	defer w.lock.Unlock()

	for i, heading := range w.successRes.Headings {
		if heading.TagName == tag {
			w.successRes.Headings[i].Levels = append(heading.Levels, level)
			return
		}
	}

	w.successRes.Headings = append(w.successRes.Headings, Heading{
		TagName: tag,
		Levels:  []string{level},
	})
}

func (w *AnalysisSuccessResponseManager) AddUrlInfo(url string, urlType int, status int, latency int64) {
	urlTypeStr := "External"
	if urlType == 0 {
		urlTypeStr = "Internal"
	}

	w.lock.Lock()
	defer w.lock.Unlock()

	w.successRes.Urls = append(w.successRes.Urls, Url{
		Url:     url,
		Type:    urlTypeStr,
		Status:  status,
		Latency: latency,
	})
}

func (w *AnalysisSuccessResponseManager) SetHasLogin(hasLogin bool) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.HasLogin = hasLogin
}

func (w *AnalysisSuccessResponseManager) SetLinkCounts(internalCount, externalCount, inaccessibleCount int) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.InternalLinkCount = internalCount
	w.successRes.ExternalLinkCount = externalCount
	w.successRes.InaccessibleLinkCount = inaccessibleCount
}

func (w *AnalysisSuccessResponseManager) ToString() string {
	w.lock.RLock()
	defer w.lock.RUnlock()

	b, err := json.Marshal(w.successRes)
	if err != nil {
		log.Println("Error occurred when marshalling the response", err)
	}
	return string(b)
}

func (w *AnalysisSuccessResponseManager) GetAnalysisResponse() AnalysisSuccessResponse {
	return w.successRes
}
