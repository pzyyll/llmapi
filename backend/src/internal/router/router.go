package router

import (
	"llmapi/src/internal/router/dashboard"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	// Set up the router for the application
	engine.GET("/ping", Ping)

	dashboard.SetupRouter(engine)
}
