package api

import (
	"net/http"

	"llmapi/src/internal/router/dashboard/api/dto"
	"llmapi/src/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService service.UserService
	authService service.AuthService
}

func NewAuthHander(userService service.UserService, authService service.AuthService) AuthHandler {
	return AuthHandler{
		userService: userService,
		authService: authService,
	}
}

func (r *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
		})
		return
	}

	ret, err := r.authService.VerifyUser(c, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		AccessToken:  ret.AccessToken,
		RefreshToken: ret.RefreshToken,
		UserID:       ret.User.ID,
		Username:     ret.User.Username,
		Email:        *ret.User.Email,
		Role:         ret.User.Role,
	})
}

func (r *AuthHandler) RefreshToken(c *gin.Context) {
	// Get Authorization header in "Bearer <token>" format
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Authorization header not found",
		})
		return
	}
	// Parse header info "Bearer" and "<token>"  AI!
	
}
