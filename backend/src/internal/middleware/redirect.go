package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RedirectToV1Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Redirect to v1
		path := c.Request.URL.Path

		if strings.HasPrefix(path, "/api") && !strings.HasPrefix(path, "/api/v1") {
			// Redirect to v1
			c.Redirect(http.StatusMovedPermanently, "/api/v1"+path[len("/api"):])
			c.Abort()
			return
		}
	}
}
