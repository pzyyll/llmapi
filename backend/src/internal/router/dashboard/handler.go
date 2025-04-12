package dashboard

import (
	"embed"
	"net/http"
	"net/url"
	"os"

	"llmapi/src/internal/middleware"
	"llmapi/src/internal/router/dashboard/api"
	"llmapi/src/pkg/logger"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const BasePrefix = "/dashboard"

//go:embed all:static
var StaticDistFS embed.FS

// SetupRouter configures all dashboard routes including API and web UI
func SetupRouter(engine *gin.Engine) {
	// Configure CORS middleware
	engine.Use(middleware.CORS())

	// Setup API routes
	setupAPIRoutes(engine)

	// Setup frontend routes - either dev redirect or static files
	setupFrontendRoutes(engine)
}

// setupAPIRoutes configures all dashboard API endpoints
func setupAPIRoutes(engine *gin.Engine) {
	apiGroup := engine.Group(BasePrefix+"/api", gzip.Gzip(gzip.DefaultCompression))
	{
		apiGroup.POST("/login", api.Login)
		// Additional API routes can be added here
	}
}

// setupFrontendRoutes configures the dashboard UI routes
// If DEV_FRONTEND_URL is set, it redirects to that URL
// Otherwise, it serves embedded static files
func setupFrontendRoutes(engine *gin.Engine) {
	devFrontendUrl := os.Getenv("DEV_FRONTEND_URL")
	if devFrontendUrl != "" {
		setupDevRedirect(engine, devFrontendUrl)
	} else {
		serveStaticFiles(engine)
	}
}

// setupDevRedirect configures redirection to dev frontend server
func setupDevRedirect(engine *gin.Engine, devFrontendUrl string) {
	baseUrl, err := url.Parse(devFrontendUrl)
	if err != nil {
		logger.Sys().Error("The frontend url at `DEV_FRONTEND_URL` is invalid.", "url", devFrontendUrl)
		return
	}

	engine.GET(BasePrefix, func(c *gin.Context) {
		targetUrl := baseUrl.JoinPath(c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, targetUrl.String())
	})
}

// serveStaticFiles serves the embedded static files and configures SPA behavior
func serveStaticFiles(engine *gin.Engine) {
	fs, err := static.EmbedFolder(StaticDistFS, "static/dist")
	if err != nil {
		logger.Sys().Error("Failed to create embed folder.", "error", err.Error())
		return
	}

	// Configure compression for static files
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	// Optionally enable these middleware
	// engine.Use(middleware.GlobalWebRateLimit())
	// engine.Use(middleware.Cache())

	// Serve static files
	engine.Use(static.Serve(BasePrefix, fs))

	// Configure SPA routing - always return index.html for unmatched routes
	setupSPARouting(engine)
}

// setupSPARouting ensures all routes in SPA return the index.html file
func setupSPARouting(engine *gin.Engine) {
	indexHtml, err := StaticDistFS.ReadFile("static/dist/index.html")
	if err != nil {
		logger.Sys().Error("Failed to read index.html.", "error", err.Error())
		return
	}

	engine.NoRoute(func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache")
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHtml)
	})
}
