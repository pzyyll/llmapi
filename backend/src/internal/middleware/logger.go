package middleware

import (
	"time"

	"llmapi/src/pkg/logger"

	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Log the request details
		duration := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		logger.Log.Info(
			"Request msg",
			map[string]interface{}{
				"status_code": statusCode,
				"client_ip":   clientIP,
				"method":      method,
				"path":        path,
				"duration":    duration,
				"type":        "request",
			},
		)
	}
}
