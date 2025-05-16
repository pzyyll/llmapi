package middleware

import (
	"fmt"

	"llmapi/src/internal/constants"
	dto "llmapi/src/internal/dto/v1"
	"llmapi/src/internal/model"
	"llmapi/src/pkg/auth"

	"github.com/gin-gonic/gin"
)

func AbortWithError(c *gin.Context, code int, error string) {
	c.JSON(code, dto.ErrorResponse{
		Code:  code,
		Error: error,
	})
	c.Abort()
}

func GetUser(c *gin.Context) (*model.User, error) {
	userVal, ok := c.Get(constants.ContextUserKey)
	if !ok {
		return nil, fmt.Errorf("user not found in context")
	}

	user, ok := userVal.(*model.User)
	if !ok {
		return nil, fmt.Errorf("invalid user type in context")
	}
	return user, nil
}

func GetAPIAuthToken(c *gin.Context) (string, error) {
	token, err := auth.GetAuthorizationTokenFromHeader(c.GetHeader(string(constants.AuthorizationHeader)), constants.AuthTypeBearer)
	if err != nil {
		return "", err
	}
	return token, nil
}
