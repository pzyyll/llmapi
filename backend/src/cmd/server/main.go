package main

import (
	"github.com/joho/godotenv"

	"llmapi/src/internal/core"
	"llmapi/src/pkg/logger"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		logger.Log.SysError("Error loading .env file", map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	
	logger.Log.SysInfo("Starting LLMAPI server", map[string]interface{}{
		"version": "1.0.0",
	})

	core.Run()
}
