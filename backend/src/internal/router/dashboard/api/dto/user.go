package dto

// User info for login
type UserInfo struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// User info for login response
type UserInfoResponse struct {
	Token    string `json:"token"`
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar"`
}

type LoginRequest struct {
	UserInfo
}
