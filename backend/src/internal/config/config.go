package config

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"llmapi/src/internal/constants"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Port               int      `koanf:"port" validate:"min=1,max=65535"`
	Host               string   `koanf:"host" validate:"hostname|ip"`
	DSN                string   `koanf:"dsn"`
	RedisURL           string   `koanf:"redis_url"`
	AccessTokenExpiry  int      `koanf:"access_token_expiry" validate:"min=1"`
	RefreshTokenExpiry int      `koanf:"refresh_token_expiry" validate:"min=1"`
	JwtSecret          string   `koanf:"jwt_secret"`
	JwtIssuer          string   `koanf:"jwt_issuer"`
	JwtSignedMethod    string   `koanf:"jwt_signed_method" validate:"oneof=HS256 HS384 HS512"`
	LogLevel           string   `koanf:"log_level" validate:"oneof=debug info warn error"`
	DBLogLevel         int      `koanf:"db_log_level" validate:"min=1,max=4"` // 1: Silent, 2: Error, 3: Warn, 4: Info
	DBAutoMigrate      bool     `koanf:"db_auto_migrate"`
	WorkerID           int64    `koanf:"worker_id"` // Unique identifier for the worker, used in distributed systems
	AdminUser          string   `koanf:"admin_user"`
	AdminPassword      string   `koanf:"admin_password"`
	AllowOrigins       []string `koanf:"allow_origins"` // Comma-separated list of allowed origins for CORS
	TurnstileSecretKey string   `koanf:"turnstile_secret_key"`
	TurnstileEnabled   bool     `koanf:"turnstile_enabled"`
	TurnstileVerifyEndpoint  string   `koanf:"turnstile_verify_endpoint"`
}

func initDefaultConfig() *Config {
	return &Config{
		Port:               13140,
		Host:               "localhost",
		DSN:                "sqlite://llmapi.db",
		RedisURL:           "",
		JwtSecret:          "",
		AccessTokenExpiry:  3600,
		RefreshTokenExpiry: 30 * 24 * 60 * 60, // 30 days
		JwtIssuer:          "llmapi",
		JwtSignedMethod:    "HS256",
		LogLevel:           "info",
		DBLogLevel:         1,
		DBAutoMigrate:      true,
		WorkerID:           1,
		AdminUser:          "admin",
		AdminPassword:      "zaq12wsx@0",
		AllowOrigins:       []string{"*"}, // Default to allow all origins
		TurnstileSecretKey: "",
		TurnstileEnabled:   false,
		TurnstileVerifyEndpoint:  "https://challenges.cloudflare.com/turnstile/v0/siteverify",
	}
}

func generateRandomSecret(length int) string {
	keyBytes := make([]byte, length)
	rand.Read(keyBytes)
	return base64.StdEncoding.EncodeToString(keyBytes)
}

func LoadConfig() *Config {
	c := initDefaultConfig()
	loadFromEnvironment(c)
	ensureJwtSecret(c)
	err := validateConfig(c)
	if err != nil {
		panic(err)
	}
	return c
}

func loadFromEnvironment(c *Config) {
	k := koanf.New("_")
	err := k.Load(env.Provider(constants.EnvPrefix, "_", func(key string) string {
		return strings.ToLower(
			strings.TrimPrefix(key, constants.EnvPrefix),
		)
	}), nil)
	if err != nil {
		panic(err)
	}
	k.UnmarshalWithConf("", c, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true})
}

func ensureJwtSecret(c *Config) {
	if c.JwtSecret != "" {
		return
	}

	data, err := os.ReadFile(constants.SecretFilePath)
	if err == nil {
		c.JwtSecret = string(data)
		return
	}

	c.JwtSecret = generateRandomSecret(constants.DesiredKeyLength)
	err = os.WriteFile(constants.SecretFilePath, []byte(c.JwtSecret), 0o644)
	if err != nil {
		panic(err)
	}

	fmt.Println("JWT secret written to:", constants.SecretFilePath)
}

func validateConfig(c *Config) error {
	if err := validator.New().Struct(c); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMsgs []string
		for _, fe := range validationErrors {
			errorMsgs = append(errorMsgs, fmt.Sprintf("'%s': require '%s %s' (value: '%v')", fe.Namespace(), fe.Tag(), fe.Param(), fe.Value()))
		}
		return fmt.Errorf("configuration validation failed:\n - %s", strings.Join(errorMsgs, "\n - "))
	}
	return nil
}
