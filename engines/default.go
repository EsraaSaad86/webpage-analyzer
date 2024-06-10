package engines

import (
	"webpage-analyzer/controllers"
	"webpage-analyzer/middleware"

	"github.com/gin-gonic/gin"
)

type DefaultEngine struct {
	analyzeController *controllers.WebPageAnalyzerController
}

func NewDefaultEngine(analyzeController *controllers.WebPageAnalyzerController) *DefaultEngine {
	return &DefaultEngine{analyzeController: analyzeController}
}

func (de *DefaultEngine) SetupRoutes() *gin.Engine {
	router := gin.Default()
	// Disable trusting all proxies
	router.SetTrustedProxies(nil)
	router.Use(middleware.CORSMiddleware())
	router.StaticFile("/", "./frontend/index.html")
	router.Static("/static", "./frontend")

	// Define routes directly without using a group
	router.GET("/v1/analyze", de.analyzeController.AnalyzeWebPage)
	return router
}
