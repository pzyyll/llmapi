package middleware

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"llmapi/src/internal/config"
	"llmapi/src/internal/constants"
	"llmapi/src/internal/utils/log"

	"github.com/gin-gonic/gin"
)

type turnstileResponse struct {
	Success    bool     `json:"success"`
	ErrorCodes []string `json:"error-codes,omitempty"`
}

func TurnstileMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := log.GetContextLogger(c)
		if !cfg.TurnstileEnabled {
			c.Next()
			return
		}

		log.Info("Turnstile start verifying token")

		secretKey := cfg.TurnstileSecretKey

		if secretKey == "" {
			log.Error("Turnstile secret key is not set")
			AbortWithError(c, http.StatusInternalServerError, "Turnstile secret key is not set")
			return
		}

		turnstileToken := c.GetHeader(string(constants.TurnstileTokenHeader))
		if turnstileToken == "" {
			AbortWithError(c, http.StatusBadRequest, "Turnstile token is required")
			return
		}

		// Verify the Turnstile token with the secret key
		client := &http.Client{Timeout: 5 * time.Second}
		data, err := client.PostForm(
			cfg.TurnstileVerifyEndpoint,
			url.Values{
				"secret":   {secretKey},
				"response": {turnstileToken},
				"remoteip": {c.ClientIP()},
			},
		)
		if err != nil {
			log.Error("Failed to verify Turnstile token", "error", err)
			AbortWithError(c, http.StatusInternalServerError, "Failed to verify Turnstile token")
			return
		}
		defer data.Body.Close()
		var response turnstileResponse
		if err := json.NewDecoder(data.Body).Decode(&response); err != nil {
			log.Error("Failed to decode Turnstile response", "error", err)
			AbortWithError(c, http.StatusInternalServerError, "Failed to decode Turnstile response")
			return
		}

		if !response.Success {
			log.Error("Turnstile verification failed", "token", turnstileToken, "ip", c.ClientIP(), "errors", response.ErrorCodes)
			AbortWithError(c, http.StatusUnauthorized, "Turnstile verification failed")
			return
		}
		log.Info("Turnstile verification succeeded")
		c.Next()
	}
}
