package middleware

import (
	"net/http"
	"strings"

	"llmapi/src/internal/constants"
	dto "llmapi/src/internal/dto/v1"
	"llmapi/src/internal/service"
	"llmapi/src/pkg/auth"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authService service.AuthService
}

func NewAuthMiddleware(authService service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// AuthMiddleware is a middleware that checks if the user is authenticated
func (a *AuthMiddleware) AccessTokenMiddleware() gin.HandlerFunc {
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

		user, err := a.authService.VerifyAccessToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:  http.StatusUnauthorized,
				Error: "Invalid access token",
			})
			c.Abort()
			return
		}

		c.Set(constants.ContextUserKey, user)

		c.Next()
	}
}

func (a *AuthMiddleware) RefreshTokenMiddleware() gin.HandlerFunc {
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

		user, err := a.authService.VerifyRefreshToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:  http.StatusUnauthorized,
				Error: "Invalid refresh token",
			})
			c.Abort()
			return
		}

		c.Set(constants.ContextUserKey, user)
		c.Set(constants.ContextRefreshTokenKey, token)

		c.Next()
	}
}
