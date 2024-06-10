package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"webpage-analyzer/analyzers"
	"webpage-analyzer/analyzers/models"
	"webpage-analyzer/http/responses"
	"webpage-analyzer/utils"

	"github.com/gin-gonic/gin"
)

type WebPageAnalyzerController struct {
	analyzers []analyzers.Analyzer
}

func NewWebPageAnalyzerController(analyzers ...analyzers.Analyzer) *WebPageAnalyzerController {
	return &WebPageAnalyzerController{analyzers: analyzers}
}

func (wpa *WebPageAnalyzerController) AnalyzeWebPage(ginCtx *gin.Context) {
	urlParam := strings.TrimSpace(ginCtx.Query("url"))
	if urlParam == "" {
		ginCtx.JSON(http.StatusBadRequest, responses.NewErrorResponse("URL parameter is required", nil))
		return
	}

	// Check if the URL is accessible
	if !utils.IsURLAccessible(urlParam) {
		ginCtx.JSON(http.StatusBadRequest, responses.NewErrorResponse("URL is not accessible", nil))
		return
	}

	resManager := responses.NewAnalysisSuccessResponseManager()

	res, err := wpa.Analyze(urlParam, resManager)
	if err != nil {
		log.Printf("Error in analyzing the webpage: %v", err)
		ginCtx.JSON(http.StatusInternalServerError, responses.NewErrorResponse("Error in analyzing the webpage", err))
		return
	}

	log.Println("All analyzers completed")
	ginCtx.JSON(http.StatusOK, res)
}

func (wpa *WebPageAnalyzerController) extractBodyFromURL(url string) (host, bodyContent string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	return resp.Request.URL.Host, string(bodyBytes), nil
}

func (wpa *WebPageAnalyzerController) Analyze(urlParam string, resManager responses.WebPageAnalyzerResponseManager) (response responses.AnalysisSuccessResponse, err error) {
	log.Println("Analysis started")

	host, body, err := wpa.extractBodyFromURL(urlParam)
	if err != nil {
		return responses.AnalysisSuccessResponse{}, err
	}

	analyzerInfo := models.NewAnalyzerInfo(body, host)

	var wg sync.WaitGroup
	wg.Add(len(wpa.analyzers))

	for _, analyzer := range wpa.analyzers {
		go func(a analyzers.Analyzer) {
			defer wg.Done()
			a.Analyze(analyzerInfo, resManager)
		}(analyzer)
	}

	wg.Wait()

	log.Println("Analysis completed")
	return resManager.GetAnalysisResponse(), nil
}
