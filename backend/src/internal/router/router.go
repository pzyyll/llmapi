package router

import (
	// "fmt"
	"net/http"
	"net/url"
	"os"

	// "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	admin "llmapi/src/internal/router/admin"
	"llmapi/src/static"
)

func SetupRouter(engine *gin.Engine) {
	// Set up the router for the application
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())


	devFrontendUrl := os.Getenv("DEV_FRONTEND_URL")
	if devFrontendUrl != "" {
		baseUrl, err := url.Parse(devFrontendUrl)
		if err != nil {
			panic(err)
		}
		// If DEV_FRONTEND_URL is set, use it as the frontend URL
		engine.GET(static.BasePrefix, func(c *gin.Context) {
			targetUrl := baseUrl.JoinPath(c.Request.URL.Path)
			c.Redirect(http.StatusMovedPermanently, targetUrl.String())
		})
	} else {
		engine.Use(static.ServeSPA())
		// 返回 index.html 文件，由客户端处理路由
		engine.NoRoute(func(c *gin.Context) {
			c.Header("Cached-Control", "no-cache")
			c.Data(http.StatusOK, "text/html; charset=utf-8", static.IndexHTML)
		})
	}


	adminGroup := engine.Group("/admin")
	{
		adminGroup.POST("/login", admin.Login)
	}
}
