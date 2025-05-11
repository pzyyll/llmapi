package dashboard

import (
	"net/http"

	"llmapi/src/internal/constants"
	"llmapi/src/internal/middleware"
	"llmapi/src/internal/model"
	"llmapi/src/internal/service"

	dto "llmapi/src/internal/dto/v1"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService service.UserService
	authService service.AuthService
}

func NewAuthHandler(userService service.UserService, authService service.AuthService) AuthHandler {
	return AuthHandler{
		userService: userService,
		authService: authService,
	}
}

func (r *AuthHandler) Login(c *gin.Context) {
	logger := middleware.GetContextLogger(c)
	logger.Info("Login request received")
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Login request binding failed", "error", err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "Bad Request",
		})
		return
	}

	ret, err := r.authService.VerifyUser(c, req.Username, req.Password)
	if err != nil {
		logger.Error("Login failed", "error", err)
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:  http.StatusUnauthorized,
			Error: err.Error(),
		})
		return
	}

	email := ""
	if ret.User.Email != nil {
		email = *ret.User.Email
	}

	logger.Debug("Login successful", "user_id", ret.User.UserID, "username", ret.User.Username)
	c.JSON(http.StatusOK, dto.LoginResponse{
		AccessToken: ret.AccessToken,
		// RefreshToken: ret.RefreshToken,
		UserID:   uint(ret.User.UserID),
		Username: ret.User.Username,
		Email:    email,
		Role:     ret.User.Role,
	})
}

func (r *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "Bad Request",
		})
		return
	}

	user, err := r.userService.CreateUser(req.Username, req.Password, constants.RoleTypeUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}

	token, refreshToken, err := r.authService.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
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
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}

	r.authService.DeleteRefreshToken(oldRefreshToken)

	c.JSON(http.StatusOK, dto.RefreshTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	})
}
