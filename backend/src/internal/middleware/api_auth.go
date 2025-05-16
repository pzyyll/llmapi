package middleware

import (
	"net/http"

	"llmapi/src/internal/constants"
	dto "llmapi/src/internal/dto/v1"
	"llmapi/src/internal/service"

	"github.com/gin-gonic/gin"
)

func APIAuthMiddleware(apiKeyService service.APIKeyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the API key from the request header
		token, err := GetAPIAuthToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:  http.StatusUnauthorized,
				Error: err.Error(),
			})
			c.Abort()
			return
		}

		// Validate the API key
		apiKeyRecord, apiUser, err := apiKeyService.ValidateAPIKey(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:  http.StatusUnauthorized,
				Error: err.Error(),
			})
			c.Abort()
			return
		}

		c.Set(constants.ContextUserKey, apiUser)
		c.Set(constants.ContentAPIRecordKey, apiKeyRecord)
		c.Next()
	}
}
