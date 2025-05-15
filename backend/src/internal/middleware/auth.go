package middleware

import (
	"net/http"
	"slices"

	"llmapi/src/internal/constants"
	dto "llmapi/src/internal/dto/v1"
	"llmapi/src/internal/model"
	"llmapi/src/internal/service"

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
		user, err := a.authService.VerifyCtxAccessToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:  http.StatusUnauthorized,
				Error: "Invalid access token: " + err.Error(),
			})
			c.Abort()
			return
		}

		c.Set(constants.ContextUserKey, user)

		c.Next()
	}
}

func (a *AuthMiddleware) AdminMiddleware(checkRoles ...constants.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Get(constants.ContextUserKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:  http.StatusUnauthorized,
				Error: "User not found in context",
			})
			c.Abort()
			return
		}

		role := constants.RoleType(user.(*model.User).Role)

		allowed := slices.Contains(checkRoles, role)
		if !allowed {
			c.JSON(http.StatusForbidden, dto.ErrorResponse{
				Code:  http.StatusForbidden,
				Error: "Permission denied",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
