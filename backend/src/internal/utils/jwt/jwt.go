package jwt

import (
	"fmt"
	"time"

	"llmapi/src/internal/constants"

	"github.com/golang-jwt/jwt/v5"
)

type LoginPayload struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Jti      int    `json:"jti,omitempty"`
	jwt.RegisteredClaims
}

func NewLoginPayload(userID uint, username string, issuer string, expire uint, jti int) *LoginPayload {
	return &LoginPayload{
		UserID:   userID,
		Username: username,
		Jti:      jti,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire) * time.Second)),
		},
	}
}

// GenerateToken creates a signed JWT string using the provided secret key and payload.
// It uses HS256 as the signing method.
func GenerateToken(signedMethod string, secretKey string, payload *LoginPayload) (string, error) {
	method := jwt.SigningMethodHS256
	switch signedMethod {
	case constants.JWT_SIGNED_HS256:
		method = jwt.SigningMethodHS256
	case constants.JWT_SIGNED_HS384:
		method = jwt.SigningMethodHS384
	case constants.JWT_SIGNED_HS512:
		method = jwt.SigningMethodHS512
	default:
		return "", fmt.Errorf("unsupported signing method: %s", signedMethod)
	}
	token := jwt.NewWithClaims(method, payload)
	return token.SignedString([]byte(secretKey))
}
