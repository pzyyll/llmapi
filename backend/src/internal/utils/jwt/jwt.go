package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// signingMethod holds the JWT signing method.
// Define it as a variable for potential future configuration or modification.
var signingMethod = jwt.SigningMethodHS256

// LoginPayload defines the structure of the JWT claims for user login.
type LoginPayload struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	// Jti (JWT ID) claim provides a unique identifier for the JWT.
	// Can be used to prevent JWT replays (once per token).
	// Also useful for correlating refresh tokens or enabling token revocation.
	Jti int64 `json:"jti,omitempty"` // Renamed field jti to Jti and changed type to int64 for compatibility with uid generator
	jwt.RegisteredClaims
}

// NewLoginPayload creates a new instance of LoginPayload with standard and custom claims.
func NewLoginPayload(userID uint, username string, issuer string, expire uint, jti int64) *LoginPayload { // Updated jti type to int64
	return &LoginPayload{
		UserID:   userID,
		Username: username,
		Jti:      jti, // Updated field name to Jti
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire) * time.Second)),
		},
	}
}

// GenerateToken creates a signed JWT string using the provided secret key and payload.
// It uses the package-level signingMethod variable.
func GenerateToken(secretKey string, payload *LoginPayload) (string, error) {
	// Use the package-level signingMethod variable
	token := jwt.NewWithClaims(signingMethod, payload)
	return token.SignedString([]byte(secretKey))
}
