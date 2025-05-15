package dashboard

import (
	"net/http"

	"llmapi/src/internal/constants"
	dto "llmapi/src/internal/dto/v1"
	"llmapi/src/internal/model"
	"llmapi/src/internal/service"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related API requests
type UserHandler struct {
	// UserService is the service for user-related operations
	UserService service.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// GetUserInfo handles the request to get user information
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	// AuthAccessTokenMiddleware is a middleware that checks if the user is authenticated
	// and sets the user in the context
	// This middleware should be applied to the route in the router
	user, _ := c.Get(constants.ContextUserKey)
	userInfo, ok := user.(*model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: "Failed to get user from context",
		})
		return
	}

	c.JSON(http.StatusOK, dto.NewUser(userInfo))
}

// UpdateUserInfo handles the request to update user information
func (h *UserHandler) UpdateUserInfo(c *gin.Context) {
	// TODO: Implement the update user information logic
	c.JSON(http.StatusOK, gin.H{
		"message": "unimplemented",
	})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: "Failed to get users: " + err.Error(),
		})
		return
	}

	retUsers := make([]dto.UserProfile, len(*users))
	for i, user := range *users {
		retUsers[i] = *dto.NewUser(&user)
	}

	c.JSON(http.StatusOK, dto.Users{
		Users: retUsers,
	})
}
