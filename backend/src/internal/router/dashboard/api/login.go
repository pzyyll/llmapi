package api

import (
	"net/http"

	"llmapi/src/internal/router/dashboard/api/dto"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Handle login logic here
	// For example, validate user credentials and set session
	// @TODO: Implement actual login logic

	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, dto.UserInfoResponse{
		Token:    "example_token",
		UserID:   1,
		Username: req.Username,
		Email:    "example@example.com",
		Role:     "user",
		Avatar:   "example_avatar.png",
	})
}
