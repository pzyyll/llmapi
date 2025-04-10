package core

import (
	"os"

	"github.com/gin-gonic/gin"

	"llmapi/src/internal/middleware"
	"llmapi/src/internal/router"
)

func Run() {
	engine := gin.New()

	// Set middleware for the engine
	engine.Use(middleware.GinLogger())
	engine.Use(gin.Recovery())
	// engine.Use(cors.Default())

	router.SetupRouter(engine)

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "13140"
	}

	address := host + ":" + port
	engine.Run(address)
}
