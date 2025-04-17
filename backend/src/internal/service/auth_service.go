package service

import (
	"context"
	"sync"

	"llmapi/src/internal/config"
	"llmapi/src/internal/constants"
	"llmapi/src/internal/model"
	"llmapi/src/internal/utils"
	"llmapi/src/internal/utils/jwt"
	"llmapi/src/pkg/auth"
)

type Result struct {
	AccessToken  string
	RefreshToken string
	User         *model.User
}

type AuthService interface {
	VerifyUser(ctx context.Context, username string, password string) (ret *Result, err error)
	VerifyRefreshToken(ctx context.Context) (ret *Result, err error)
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

func (s *authService) VerifyUser(ctx context.Context, username string, password string) (ret *Result, err error) {
	user, err := s.userService.GetUserByName(username)
	if err != nil {
		return nil, err
	}

	if err := auth.CheckPasswordHash(password, user.Password); err != nil {
		return nil, err
	}

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
		return nil, err
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
		return nil, err
	}
	// Store the refresh token
	refreshTokenMap, _ := s.refreshTokenCache.LoadOrStore(user.UserID, &sync.Map{})
	refreshTokenMap.(*sync.Map).Store(tokenid, true)

	return &Result{
		AccessToken:  token,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

func (s *authService) VerifyRefreshToken(ctx context.Context) (ret *Result, err error) {
	// Parse the refresh token
	// authHeader := ctx.GetHeader("Authorization")
	return nil, nil
}