package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginPayloayd struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Jti      int    `json:"jti,omitempty"`
	jwt.RegisteredClaims
}

func NewLoginPayload(userID uint, username string, issuer string, expire uint, jti int) *LoginPayloayd {
	return &LoginPayloayd{
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
func GenerateToken(secretKey string, payload *LoginPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(secretKey))
}
