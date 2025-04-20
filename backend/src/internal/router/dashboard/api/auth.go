package api

import (
	"net/http"

	"llmapi/src/internal/constants"
	"llmapi/src/internal/model"
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

	email := ""
	if ret.User.Email != nil {
		email = *ret.User.Email
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		AccessToken:  ret.AccessToken,
		RefreshToken: ret.RefreshToken,
		UserID:       ret.User.ID,
		Username:     ret.User.Username,
		Email:        email,
		Role:         ret.User.Role,
	})
}

func (r *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
		})
		return
	}

	user, err := r.userService.CreateUser(req.Username, req.Password, constants.RoleTypeUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	token, refreshToken, err := r.authService.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.RegisterResponse{
		UserID:       user.ID,
		Username:     user.Username,
		Role:         user.Role,
		AccessToken:  token,
		RefreshToken: refreshToken,
	})
}

func (r *AuthHandler) RefreshToken(c *gin.Context) {
	user := c.MustGet(constants.ContextUserKey).(*model.User)
	oldRefreshToken := c.MustGet(constants.ContextRefreshTokenKey).(string)

	newAccessToken, newRefreshToken, err := r.authService.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			})
			return
	}

	r.authService.DeleteRefreshToken(oldRefreshToken)

	c.JSON(http.StatusOK, dto.RefreshTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	})
}
