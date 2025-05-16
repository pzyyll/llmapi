package dashboard

import (
	"net/http"
	"strconv"

	"llmapi/src/internal/constants"
	dto "llmapi/src/internal/dto/v1"
	"llmapi/src/internal/middleware"
	"llmapi/src/internal/model"
	"llmapi/src/internal/service"
	"llmapi/src/internal/utils/log"
	"llmapi/src/internal/utils/role"

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

func (h *UserHandler) DeleteUser(c *gin.Context) {
	log := log.GetContextLogger(c)
	user, err := middleware.GetUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		c.Abort()
		return
	}

	if !role.IsAdmin(constants.RoleType(user.Role)) {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{
			Code:  http.StatusForbidden,
			Error: "Permission denied",
		})
		c.Abort()
		return
	}

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		c.Abort()
		return
	}

	targetUser, err := h.UserService.GetUserByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		c.Abort()
		return
	}

	roleLevel := role.GetRoleLevel(constants.RoleType(targetUser.Role))
	if roleLevel >= role.GetRoleLevel(constants.RoleType(user.Role)) {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{
			Code:  http.StatusForbidden,
			Error: "Permission denied",
		})
		c.Abort()
		return
	}

	log.Info("Deleting user", "user_id", userID)
	err = h.UserService.DeleteUser(targetUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User deleted successfully",
	})
	c.Abort()
}
