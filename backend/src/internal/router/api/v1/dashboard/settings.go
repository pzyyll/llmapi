package dashboard

import (
	"net/http"

	"llmapi/src/internal/config"
	dto "llmapi/src/internal/dto/v1"

	"github.com/gin-gonic/gin"
)

func LoadSettingsHandler(cfg *config.Config) gin.HandlerFunc {
	// Get the settings from the request
	return func(c *gin.Context) {
		turnstileSiteKey := ""
		if cfg.TurnstileEnabled {
			turnstileSiteKey = cfg.TurnstileSiteKey
		}
		loadSettingsResponse := dto.LoadSettingsResponse{
			TurnstileSiteKey: turnstileSiteKey,
		}
		c.JSON(http.StatusOK, loadSettingsResponse)
	}
}
