package middleware

import (
	"net/http"
	"strings"

	"llmapi/src/internal/constants"
	"llmapi/src/internal/dto"
	"llmapi/src/pkg/auth"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// AuthMiddleware is a middleware that checks if the user is authenticated
func (a *AuthMiddleware) AuthAccessTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is authenticated
		protocol, token, err := auth.GetAuthorizationToken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:  http.StatusUnauthorized,
				Error: "Authorization header not found",
			})
			c.Abort()
			return
		}

		if !strings.EqualFold(protocol, constants.AuthTypeBearer) {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:  http.StatusUnauthorized,
				Error: "Authorization type not supported",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
