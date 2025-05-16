package v1

import (
	"strconv"
	"time"

	"llmapi/src/internal/model"
)

type APIKeyProfile struct {
	UserID      string     `json:"user_id"`
	Name        string     `json:"name"`
	Scopes      int64      `json:"scopes"`
	LookupKey   string     `json:"lookup_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	LastUsedAt  *time.Time `json:"last_used_at"`
	ExpireAt    *time.Time `json:"expire_at"`
	SecretBrief string     `json:"secret_brief"`
}

type CreateApiKeyRequest struct {
	// Name of the API key
	Name string `json:"name" binding:"required"`
	// Expiration time in seconds, 0 means never expire
	Expire int64 `json:"expire"`
	// Option: Scopes
	Scopes []string `json:"scopes"`
}

type CreateApiKeyResponse struct {
	// API key Info
	ApiKey APIKeyProfile `json:"key"`
	// API key secret
	Secret string `json:"secret"`
}

type GetApiKeysResponse struct {
	// API keys
	APIKeys []APIKeyProfile `json:"api_keys"`
}

func ToAPIKeyProfile(apiKey *model.APIKeyRecord) *APIKeyProfile {
	return &APIKeyProfile{
		UserID:      strconv.FormatInt(apiKey.UserID, 10),
		Name:        apiKey.Name,
		Scopes:      apiKey.Scopes,
		LookupKey:   apiKey.LookupKey,
		CreatedAt:   apiKey.CreatedAt,
		UpdatedAt:   apiKey.UpdatedAt,
		LastUsedAt:  apiKey.LastUsedAt,
		ExpireAt:    apiKey.ExpiresAt,
		SecretBrief: apiKey.SecretBrief,
	}
}
