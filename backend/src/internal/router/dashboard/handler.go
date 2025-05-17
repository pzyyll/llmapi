package dashboard

import (
	"embed"
	"net/http"
	"net/url"
	"os"
	"path"

	"llmapi/src/internal/config"
	"llmapi/src/internal/constants"
	"llmapi/src/internal/middleware"
	dashboardApiV1 "llmapi/src/internal/router/api/v1/dashboard"
	"llmapi/src/internal/service"
	"llmapi/src/internal/utils/log"
	"llmapi/src/internal/utils/role"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//go:embed all:static
var StaticDistFS embed.FS

type Options struct {
	Engine    *gin.Engine
	UserSvc   service.UserService
	AuthSvc   service.AuthService
	APIKeySvc service.APIKeyService
	Cfg       *config.Config
}

// SetupRouter configures all dashboard routes including API and web UI
func SetupRouter(opts *Options) {
	engine := opts.Engine

	// Configure CORS middleware
	engine.Use(middleware.CORS(opts.Cfg))

	// Setup API routes
	setupAPIRoutes(opts)

	// Setup frontend routes - either dev redirect or static files
	setupFrontendRoutes(opts)
}

// setupAPIRoutes configures all dashboard API endpoints
func setupAPIRoutes(opts *Options) {
	engine := opts.Engine

	apiGroup := engine.Group(path.Join(constants.DashboardPrefix, constants.DashboardApiPath),
		middleware.IpLimiterMiddleware(),
		gzip.Gzip(gzip.DefaultCompression))
	{
		authHandler := dashboardApiV1.NewAuthHandler(opts.UserSvc, opts.AuthSvc)

		turnstileMiddleware := middleware.TurnstileMiddleware(opts.Cfg)
		apiGroup.POST("/login", turnstileMiddleware, authHandler.Login)
		apiGroup.POST("/register", turnstileMiddleware, authHandler.Register)
		apiGroup.POST("/renew_token", authHandler.RefreshToken)
		// Additional API routes can be added here

		authMiddleware := middleware.NewAuthMiddleware(opts.AuthSvc)
		authenticatedGroup := apiGroup.Group("") // 创建子分组
		authenticatedGroup.Use(authMiddleware.AccessTokenMiddleware())
		{
			userHandler := dashboardApiV1.NewUserHandler(opts.UserSvc)
			authenticatedGroup.POST("/validate_token", authHandler.ValidateToken)
			authenticatedGroup.POST("/profile", userHandler.GetUserInfo)
			authenticatedGroup.POST("/update_profile", userHandler.UpdateUserInfo)
			authenticatedGroup.POST("/logout", authHandler.Logout)

			adminMiddleware := authMiddleware.AdminMiddleware(role.GetRoleLevel(constants.RoleTypeAdmin))
			authenticatedGroup.GET("/users", adminMiddleware, userHandler.GetUsers)
			authenticatedGroup.DELETE("/delete_user", adminMiddleware, userHandler.DeleteUser)

			apiKeyHandler := dashboardApiV1.NewApiKeyHandler(opts.APIKeySvc)
			authenticatedGroup.POST("/create_api_key", adminMiddleware, apiKeyHandler.CreateApiKey)
			authenticatedGroup.GET("/api_keys", adminMiddleware, apiKeyHandler.GetApiKeys)
			authenticatedGroup.DELETE("/delete_api_key", adminMiddleware, apiKeyHandler.DeleteApiKey)
		}
	}
}

// setupFrontendRoutes configures the dashboard UI routes
// If DEV_FRONTEND_URL is set, it redirects to that URL
// Otherwise, it serves embedded static files
func setupFrontendRoutes(opts *Options) {
	engine := opts.Engine
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
		log.Sys().Error("The frontend url at `DEV_FRONTEND_URL` is invalid.", "url", devFrontendUrl)
		return
	}

	engine.GET(constants.DashboardPrefix, func(c *gin.Context) {
		targetUrl := baseUrl.JoinPath(c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, targetUrl.String())
	})
}

// serveStaticFiles serves the embedded static files and configures SPA behavior
func serveStaticFiles(engine *gin.Engine) {
	fs, err := static.EmbedFolder(StaticDistFS, "static/dist")
	if err != nil {
		log.Sys().Error("Failed to create embed folder.", "error", err.Error())
		return
	}

	// Configure compression for static files
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	// Optionally enable these middleware
	// engine.Use(middleware.GlobalWebRateLimit())
	// engine.Use(middleware.Cache())

	// Serve static files
	engine.Use(static.Serve(constants.DashboardPrefix, fs))

	// Configure SPA routing - always return index.html for unmatched routes
	setupSPARouting(engine)
}

// setupSPARouting ensures all routes in SPA return the index.html file
func setupSPARouting(engine *gin.Engine) {
	indexHtml, err := StaticDistFS.ReadFile("static/dist/index.html")
	if err != nil {
		log.Sys().Error("Failed to read index.html.", "error", err.Error())
		return
	}

	engine.NoRoute(func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache")
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHtml)
	})
}
