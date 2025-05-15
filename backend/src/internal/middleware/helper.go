package middleware

import (
	dto "llmapi/src/internal/dto/v1"

	"github.com/gin-gonic/gin"
)

func AbortWithError(c *gin.Context, code int, error string) {
	c.JSON(code, dto.ErrorResponse{
		Code:  code,
		Error: error,
	})
	c.Abort()
}
