package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// LoginPayload defines the structure of the JWT claims for user login.
type LoginPayload struct { // Renamed LoginPayloayd to LoginPayload
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	// jti (JWT ID) claim provides a unique identifier for the JWT.
	// Can be used to prevent JWT replays (once per token).
	// Also useful for correlating refresh tokens or enabling token revocation.
	jti int `json:"jti,omitempty"` // Renamed field Jti to jti
	jwt.RegisteredClaims
}

// NewLoginPayload creates a new instance of LoginPayload with standard and custom claims.
func NewLoginPayload(userID uint, username string, issuer string, expire uint, jti int) *LoginPayload { // Updated return type
	return &LoginPayload{ // Updated struct literal type
		UserID:   userID,
		Username: username,
		jti:      jti, // Updated field name
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
