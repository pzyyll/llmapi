package middleware

import (
	"net/http"
	"path"
	"strings"

	"llmapi/src/internal/constants"

	"github.com/gin-gonic/gin"
)

func RedirectToV1Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Redirect to v1
		targetPath := c.Request.URL.Path
		apiVersion := path.Join(constants.APIPrefix, constants.APIVersion)

		if strings.HasPrefix(targetPath, constants.APIPrefix) && !strings.HasPrefix(targetPath, apiVersion) {
			// Default API Redirect to v1
			c.Redirect(http.StatusMovedPermanently, path.Join(apiVersion, targetPath[len(constants.APIPrefix):]))
			c.Abort()
			return
		}
	}
}
