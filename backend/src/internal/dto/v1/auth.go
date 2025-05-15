package v1

type ValidateTokenResponse struct {
	UserProfile `json:"user"`
}

type RefreshTokenResponse struct {
	UserProfile `json:"user"`
	AccessToken string `json:"access_token"`
}
