package config

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

const (
	secretFilePath   = ".jwt_secret.key"
	desiredKeyLength = 64
	env_prefix       = "LA_"
)

type Config struct {
	Port       int    `koanf:"port" validate:"min=1,max=65535"`
	Host       string `koanf:"host" validate:"hostname|ip"`
	DSN        string `koanf:"dsn"`
	RedisURL   string `koanf:"redis_url"`
	JwtSecret  string `koanf:"jwt_secret"`
	JwtExpiry  int    `koanf:"jwt_expiry" validate:"min=3600"`
	JwtIssuer  string `koanf:"jwt_issuer"`
	LogLevel   string `koanf:"log_level" validate:"oneof=debug info warn error"`
	DBLogLevel int    `koanf:"db_log_level" validate:"min=1,max=4"` // 1: Silent, 2: Error, 3: Warn, 4: Info
}

func initDefaultConfig() *Config {
	return &Config{
		Port:       13140,
		Host:       "localhost",
		DSN:        "sqlite://llmapi.db",
		RedisURL:   "",
		JwtSecret:  "",
		JwtExpiry:  30 * 24 * 60 * 60, // 30 days in seconds
		JwtIssuer:  "llmapi",
		LogLevel:   "info",
		DBLogLevel: 1,
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
	err := k.Load(env.Provider(env_prefix, "_", func(key string) string {
		return strings.ToLower(
			strings.TrimPrefix(key, env_prefix),
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

	data, err := os.ReadFile(secretFilePath)
	if err == nil {
		c.JwtSecret = string(data)
		return
	}

	c.JwtSecret = generateRandomSecret(desiredKeyLength)
	err = os.WriteFile(secretFilePath, []byte(c.JwtSecret), 0o644)
	if err != nil {
		panic(err)
	}
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
