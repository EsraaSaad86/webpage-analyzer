package main

import (
	"log"
	"webpage-analyzer/analyzers"
	"webpage-analyzer/configuration"
	"webpage-analyzer/controllers"
	"webpage-analyzer/engines"
)

func main() {
	// Load configuration
	config, err := configuration.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize analyzers
	htmlVersionAnalyzer := analyzers.NewHTMLVersionAnalyzer()
	titleAnalyzer := analyzers.NewPageTitleAnalyzer()
	headingAnalyzer := analyzers.NewHeaderAnalyzer()
	linkAnalyzer := analyzers.NewLinksAnalyzer()
	loginAnalyzer := analyzers.NewLoginFormAnalyzer()

	// Initialize controllers
	analyzeController := controllers.NewWebPageAnalyzerController(
		htmlVersionAnalyzer, titleAnalyzer, headingAnalyzer, linkAnalyzer, loginAnalyzer)

	// Initialize engine
	defaultEngine := engines.NewDefaultEngine(analyzeController)

	// Setup routes and start server
	router := defaultEngine.SetupRoutes()
	router.Run(":" + config.Server.Port)
}
