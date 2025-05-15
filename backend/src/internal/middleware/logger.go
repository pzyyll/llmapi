package middleware

import (
	"time"

	"llmapi/src/internal/constants"
	"llmapi/src/internal/utils"
	"llmapi/src/internal/utils/log"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	log := log.WithType(log.RequestType)
	return func(c *gin.Context) {
		start := time.Now()

		// Generate or retrieve Request ID
		requestID := c.Request.Header.Get(constants.HttpRequestIDHeader)
		if requestID == "" {
			requestID = utils.UUID()
			c.Header(constants.HttpRequestIDHeader, requestID)
		}

		reqLogger := log.With(constants.HttpRequestIDKey, requestID)
		c.Set(constants.ContextLoggerKey, reqLogger)

		// Process request
		c.Next()

		// Log the request details
		duration := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		reqLogger.Info(
			"Request msg",
			"status_code", statusCode,
			"client_ip", clientIP,
			"method", method,
			"path", path,
			"duration", duration,
		)
	}
}
