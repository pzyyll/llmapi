package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	config := LoadConfig()
	assert.NotNil(t, config)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, "sqlite://llmapi.db", config.DSN)
	assert.Equal(t, "", config.RedisURL)
	// Check if the JWT secret is generated and not ""
	assert.NotEqual(t, "", config.JwtSecret)
	assert.Equal(t, 30*24*60*60, config.JwtExpiry) // 30 days in seconds
	assert.Equal(t, "llmapi", config.JwtIssuer)
	assert.Equal(t, "info", config.LogLevel)
	// Add more tests for other fields as needed
}

func TestLoadEnv(t *testing.T) {
	os.Setenv("LA_PORT", "9090")
	os.Setenv("LA_HOST", "127.0.0.1")
	os.Setenv("LA_DSN", "postgres://user:password@localhost:5432/dbname")
	os.Setenv("LA_REDIS_URL", "redis://localhost:6379")
	os.Setenv("LA_JWT_SECRET", "mysecret")
	os.Setenv("LA_JWT_EXPIRY", "3600")
	os.Setenv("LA_JWT_ISSUER", "myissuer")
	os.Setenv("LA_LOG_LEVEL", "debug")
	os.Setenv("LA_DB_LOG_LEVEL", "5")

	config := LoadConfig()
	assert.NotNil(t, config)
	assert.Equal(t, 9090, config.Port)
	assert.Equal(t, "127.0.0.1", config.Host)
	assert.Equal(t, "postgres://user:password@localhost:5432/dbname", config.DSN)
	assert.Equal(t, "redis://localhost:6379", config.RedisURL)
	assert.Equal(t, "mysecret", config.JwtSecret)
	assert.Equal(t, 3600, config.JwtExpiry)
	assert.Equal(t, "myissuer", config.JwtIssuer)
	assert.Equal(t, "debug", config.LogLevel)
}

func TestGenerateRandomSecret(t *testing.T) {
	secret := generateRandomSecret(32)
	assert.NotEmpty(t, secret)
	assert.Len(t, secret, 44) // Base64 encoding increases the length
	// Add more assertions as needed
}