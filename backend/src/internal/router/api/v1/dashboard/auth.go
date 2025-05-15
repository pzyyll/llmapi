package dashboard

import (
	"net/http"

	"llmapi/src/internal/constants"
	"llmapi/src/internal/model"
	"llmapi/src/internal/service"
	"llmapi/src/internal/utils/log"

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
	logger := log.GetContextLogger(c)
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

	c.SetCookie(constants.CookieNameRefreshToken, ret.RefreshToken, 0, "", "", false, true)

	logger.Debug("Login successful", "user_id", ret.User.UserID, "username", ret.User.Username)
	c.JSON(http.StatusOK, dto.LoginResponse{
		AccessToken: ret.AccessToken,
		UserProfile: *dto.NewUser(ret.User),
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

	// Set refresh token to Cookie
	c.SetCookie(constants.CookieNameRefreshToken, refreshToken, 0, "", "", false, true)

	c.JSON(http.StatusOK, dto.RegisterResponse{
		UserProfile: *dto.NewUser(user),
		AccessToken: token,
	})
}

func (r *AuthHandler) RefreshToken(c *gin.Context) {
	ret, err := r.authService.RefreshToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:  http.StatusUnauthorized,
			Error: err.Error(),
		})
		return
	}

	c.SetCookie(constants.CookieNameRefreshToken, ret.RefreshToken, 0, "", "", false, true)

	c.JSON(http.StatusOK, dto.RefreshTokenResponse{
		UserProfile: *dto.NewUser(ret.User),
		AccessToken: ret.AccessToken,
	})
}

func (r *AuthHandler) ValidateToken(c *gin.Context) {
	user := c.MustGet(constants.ContextUserKey).(*model.User)

	c.JSON(http.StatusOK, dto.ValidateTokenResponse{
		UserProfile: *dto.NewUser(user),
	})
}

func (r *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie(constants.CookieNameRefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "Bad Request",
		})
		return
	}

	if err := r.authService.DeleteRefreshToken(refreshToken); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}

	c.SetCookie(constants.CookieNameRefreshToken, "", -1, "", "", false, true)

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Logout successful",
	})
}
