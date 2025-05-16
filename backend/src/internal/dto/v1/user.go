package v1

import (
	"llmapi/src/internal/model"
	"strconv"
)

// User info for login
type UserInfo struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// User info for login response
type LoginResponse struct {
	UserProfile `json:"user"`
	AccessToken string `json:"access_token"`
}

type LoginRequest struct {
	UserInfo
}

type RegisterRequest struct {
	UserInfo
}

type RegisterResponse struct {
	UserProfile `json:"user"`
	AccessToken string `json:"access_token"`
}

type DeleteUserRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
}

type UserProfile struct {
	UserID    string   `json:"user_id"`
	Username  string  `json:"username"`
	Email     *string `json:"email"`
	Role      string  `json:"role"`
	IsActive  bool    `json:"is_active"`
	CreatedAt string  `json:"created_at"`
}

type Users struct {
	Users []UserProfile `json:"users"`
}

func NewUser(user *model.User) *UserProfile {
	return &UserProfile{
		UserID:    strconv.FormatInt(user.UserID, 10),
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
