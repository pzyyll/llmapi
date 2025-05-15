package service

import (
	"fmt"
	"strings"
	"sync"

	"llmapi/src/internal/config"
	"llmapi/src/internal/constants"
	"llmapi/src/internal/model"
	"llmapi/src/internal/utils"
	"llmapi/src/internal/utils/jwt"
	"llmapi/src/internal/utils/log"
	"llmapi/src/pkg/auth"

	"github.com/gin-gonic/gin"
)

type Result struct {
	AccessToken  string
	RefreshToken string
	User         *model.User
}

type AuthService interface {
	VerifyUser(ctx *gin.Context, username string, password string) (ret *Result, err error)
	VerifyRefreshToken(token string) (user *model.User, err error)
	VerifyAccessToken(token string) (user *model.User, err error)
	VerifyCtxAccessToken(ctx *gin.Context) (*model.User, error)
	VerifyCtxRefreshToken(ctx *gin.Context) (*model.User, string, error)

	RefreshToken(ctx *gin.Context) (ret *Result, err error)
	CreateToken(user *model.User) (access_token, refresh_token string, err error)
	DeleteRefreshToken(token string) error
}

type authService struct {
	userService       UserService
	uidGenerator      utils.UidGenerator
	cfg               *config.Config
	refreshTokenCache sync.Map
}

func NewAuthService(userService UserService, cfg *config.Config, uidGenrator utils.UidGenerator) AuthService {
	return &authService{
		userService:  userService,
		uidGenerator: uidGenrator,
		cfg:          cfg,
	}
}

func (s *authService) CreateToken(user *model.User) (access_token, refresh_token string, err error) {
	// Generate access token
	payload := jwt.NewLoginPayload(
		user.ID,
		user.Username,
		constants.AppName,
		uint(s.cfg.AccessTokenExpiry),
		0,
	)
	token, err := jwt.GenerateToken(s.cfg.JwtSignedMethod, s.cfg.JwtSecret, payload)
	if err != nil {
		return "", "", err
	}

	tokenid := s.uidGenerator.GenerateUID()
	rPayload := jwt.NewLoginPayload(
		uint(user.UserID),
		user.Username,
		constants.AppName,
		uint(s.cfg.RefreshTokenExpiry),
		int(tokenid),
	)
	refreshToken, err := jwt.GenerateToken(s.cfg.JwtSignedMethod, s.cfg.JwtSecret, rPayload)
	if err != nil {
		return "", "", err
	}
	// Store the refresh token
	log.Sys().Debug("Storing refresh token in cache", "tokenid", tokenid, "user_id", user.UserID)
	s.refreshTokenCache.Store(rPayload.Jti, &sync.Map{})
	return token, refreshToken, nil
}

func (s *authService) DeleteRefreshToken(token string) error {
	// Delete the refresh token from the cache

	claims, err := jwt.ParseToken(token, s.cfg.JwtSecret)
	if err != nil {
		log.Sys().Error("Failed to parse refresh token", "error", err)
		return err
	}

	jti := claims.Jti

	if _, ok := s.refreshTokenCache.Load(jti); ok {
		s.refreshTokenCache.Delete(jti)
		log.Sys().Debug("Deleted refresh token from cache", "jti", jti)
	} else {
		log.Sys().Debug("Refresh token not found in cache", "jti", jti)
		return fmt.Errorf("refresh token not found")
	}
	return nil
}

func (s *authService) VerifyUser(ctx *gin.Context, username string, password string) (ret *Result, err error) {
	log := log.GetContextLogger(ctx)

	user, err := s.userService.GetUserByName(username)
	if err != nil {
		log.Error("Failed to get user by name", "error", err)
		return nil, err
	}

	if err := auth.CheckPasswordHash(password, user.Password); err != nil {
		log.Error("Failed to check password hash", "error", err)
		return nil, err
	}

	refreshToken, err := ctx.Cookie(constants.CookieNameRefreshToken)
	if err == nil {
		log.Info("Refresh token found in cookie")
		s.DeleteRefreshToken(refreshToken)
	}

	// Generate access token
	token, refreshToken, err := s.CreateToken(user)
	if err != nil {
		return nil, err
	}

	return &Result{
		AccessToken:  token,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

func (s *authService) VerifyAccessToken(token string) (user *model.User, err error) {
	// Parse the access token
	claims, err := jwt.ParseToken(token, s.cfg.JwtSecret)
	if err != nil {
		return nil, err
	}

	return s.userService.GetUserByID(int64(claims.UserID))
}

func (s *authService) VerifyRefreshToken(token string) (user *model.User, err error) {
	// Parse the refresh token
	// authHeader := ctx.GetHeader("Authorization")

	claims, err := jwt.ParseToken(token, s.cfg.JwtSecret)
	if err != nil {
		log.Sys().Error("Failed to parse refresh token", "error", err)
		return nil, err
	}

	if _, ok := s.refreshTokenCache.Load(claims.Jti); !ok {
		log.Sys().Error("Refresh token not found in cache", "user_id", claims.UserID, "jti", claims.Jti)
		return nil, fmt.Errorf("refresh token invalid")
	}

	// Get the user by ID
	user, err = s.userService.GetUserByUserID(int64(claims.UserID))
	if err != nil {
		log.Sys().Error("Failed to get user by ID", "user_id", claims.UserID, "error", err)
		return nil, err
	}

	return user, nil
}

func (s *authService) VerifyCtxAccessToken(ctx *gin.Context) (*model.User, error) {
	// Get the user from the context
	protocol, token, err := auth.GetAuthorizationToken(ctx.GetHeader("Authorization"))
	if err != nil {
		return nil, err
	}

	if !strings.EqualFold(protocol, constants.AuthTypeBearer) {
		return nil, fmt.Errorf("authorization type not supported")
	}

	// Verify the access token
	user, err := s.VerifyAccessToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) VerifyCtxRefreshToken(ctx *gin.Context) (*model.User, string, error) {
	log := log.GetContextLogger(ctx)
	// Get Refresh Token From Cookie
	refreshToken, err := ctx.Cookie(constants.CookieNameRefreshToken)
	if err != nil {
		log.Error("Failed to get refresh token from cookie", "error", err)
		return nil, "", err
	}

	// Verify the refresh token
	user, err := s.VerifyRefreshToken(refreshToken)
	if err != nil {
		log.Error("Failed to verify refresh token", "error", err)
		return nil, "", err
	}

	return user, refreshToken, nil
}

func (s *authService) RefreshToken(ctx *gin.Context) (ret *Result, err error) {
	// Get the refresh token from the context
	user, refreshToken, err := s.VerifyCtxRefreshToken(ctx)
	if err != nil {
		return nil, err
	}

	// Generate new access token and refresh token
	newAccessToken, newRefreshToken, err := s.CreateToken(user)
	if err != nil {
		return nil, err
	}

	// Delete the old refresh token
	s.DeleteRefreshToken(refreshToken)

	ret = &Result{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		User:         user,
	}
	return ret, nil
}
