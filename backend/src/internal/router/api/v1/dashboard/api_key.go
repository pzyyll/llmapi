package dashboard

import (
	"net/http"

	dto "llmapi/src/internal/dto/v1"
	"llmapi/src/internal/middleware"
	"llmapi/src/internal/service"

	"github.com/gin-gonic/gin"
)

type ApiKeyHandler interface {
	// CreateApiKey creates a new API key for the user
	CreateApiKey(c *gin.Context)
	// GetApiKeys retrieves all API keys for the user
	GetApiKeys(c *gin.Context)
	// DeleteApiKey deletes an API key for the user
	DeleteApiKey(c *gin.Context)
	// UpdateApiKey updates an API key for the user
	UpdateApiKey(c *gin.Context)
}

type apiKeyHandler struct {
	// userService is the service for user-related operations
	apiKeyService service.APIKeyService
}

func NewApiKeyHandler(apiKeyService service.APIKeyService) ApiKeyHandler {
	return &apiKeyHandler{
		apiKeyService: apiKeyService,
	}
}

func (h *apiKeyHandler) CreateApiKey(c *gin.Context) {
	user, err := middleware.GetUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}

	var req dto.CreateApiKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	apiKey, apiKeyRecord, err := h.apiKeyService.CreateAPIKey(user, req.Name, 0, req.Expire)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.CreateApiKeyResponse{
		ApiKey: *dto.ToAPIKeyProfile(apiKeyRecord),
		Secret: apiKey,
	})
}

func (h *apiKeyHandler) GetApiKeys(c *gin.Context) {
	user, err := middleware.GetUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}

	apiKeys, err := h.apiKeyService.GetAPIKeys(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}

	retApiKeys := make([]dto.APIKeyProfile, len(apiKeys))
	for i, apiKey := range apiKeys {
		retApiKeys[i] = *dto.ToAPIKeyProfile(apiKey)
	}

	c.JSON(http.StatusOK, dto.GetApiKeysResponse{
		APIKeys: retApiKeys,
	})
}

func (h *apiKeyHandler) DeleteApiKey(c *gin.Context) {
	lookupKey := c.Query("lookup_key")
	if lookupKey == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "lookup_key is required",
		})
		c.Abort()
		return
	}

	err := h.apiKeyService.DeleteAPIKeyRecordByLookupKey(lookupKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "API key deleted successfully",
	})
}

func (h *apiKeyHandler) UpdateApiKey(c *gin.Context) {
	lookupKey := c.Param("key")
	if lookupKey == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "lookup_key is required",
		})
		c.Abort()
		return
	}
	apiKeyRecord, err := h.apiKeyService.GetAPIKeyRecordByLookupKey(lookupKey)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:  http.StatusNotFound,
			Error: "API key not found",
		})
		return
	}

	var req dto.UpdateApiKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "Invalid request",
		})
		return
	}

	hasChanged := false

	if req.Name != "" && apiKeyRecord.Name != req.Name {
		apiKeyRecord.Name = req.Name
		hasChanged = true
	}

	if !hasChanged {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "No changes",
		})
		return
	}

	err = h.apiKeyService.UpdateAPIKeyRecord(apiKeyRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Updated successfully",
	})
}
