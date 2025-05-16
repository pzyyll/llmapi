package middleware

import (
	"net/http"

	"llmapi/src/internal/constants"
	dto "llmapi/src/internal/dto/v1"
	"llmapi/src/internal/service"
	"llmapi/src/internal/utils/role"

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

func (a *AuthMiddleware) AdminMiddleware(roleLevel constants.RoleLevel) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:  http.StatusUnauthorized,
				Error: err.Error(),
			})
			c.Abort()
			return
		}

		userRoleLevel := role.GetRoleLevel(constants.RoleType(user.Role))
		if userRoleLevel < roleLevel {
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
