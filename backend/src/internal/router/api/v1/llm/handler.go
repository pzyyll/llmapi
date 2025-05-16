package llm

import (
	"net/http"

	"llmapi/src/internal/middleware"
	"llmapi/src/internal/service"

	"github.com/gin-gonic/gin"
)

func TestPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func SetupRouter(router *gin.Engine, apiKeyService service.APIKeyService) {
	apiAuthMid := middleware.APIAuthMiddleware(apiKeyService)

	apiAuthGroup := router.Group("v1", apiAuthMid)
	{
		apiAuthGroup.Match([]string{http.MethodGet, http.MethodPost}, "/ping", TestPing)
	}
}
