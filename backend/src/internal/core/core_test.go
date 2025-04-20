package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"llmapi/src/pkg/logger"
)


func init() {
	gin.SetMode(gin.TestMode)
	InitServer()
}


func TestInitServer(t *testing.T) {
	err := InitServer()
	if err != nil {
		t.Fatalf("InitServer() failed: %v", err)
	}
}

func TestRouter(t *testing.T) {
	err := InitServer()
	if err != nil {
		t.Fatalf("InitServer() failed: %v", err)
	}
	logger.SetLevelString("debug")
	cfg.AccessTokenExpiry = 1

	if router == nil {
		t.Fatal("router is not initialized")
	}

	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())

	// Test the /dashboard/api/login endpoint
	requestBody := fmt.Sprintf(`{"username":"%s","password":"%s"}`, cfg.AdminUser, cfg.AdminPassword)
	req, err = http.NewRequest(http.MethodPost, "/dashboard/api/login", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user_id")
	assert.Contains(t, w.Body.String(), "username")
	assert.Contains(t, w.Body.String(), "role")
	assert.Contains(t, w.Body.String(), "refresh_token")
	assert.Contains(t, w.Body.String(), "access_token")

	// Get the access token from the login response
	var loginResponse map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &loginResponse); err != nil {
		t.Fatalf("Failed to unmarshal login response: %v", err)
	}
	accessToken, ok := loginResponse["access_token"].(string)
	if !ok {
		t.Fatal("Failed to get access token from login response")
	}
	refreshToken, ok := loginResponse["refresh_token"].(string)
	if !ok {
		t.Fatal("Failed to get refresh token from login response")
	}

	// Test the /dashboard/api/profile endpoint with `Authorization` header
	req, err = http.NewRequest(http.MethodPost, "/dashboard/api/profile", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user_id")
	assert.Contains(t, w.Body.String(), "username")

	req.Header.Set("Authorization", "Bearer "+refreshToken)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	
	time.Sleep(2 * time.Second)
	
	// Test the /dashboard/api/renew_token endpoint with `Authorization` header
	req, err = http.NewRequest(http.MethodPost, "/dashboard/api/renew_token", nil)
	req.Header.Set("Authorization", "Bearer "+refreshToken)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "access_token")
	assert.Contains(t, w.Body.String(), "refresh_token")

	renewedTokenResponse := make(map[string]interface{})
	if err := json.Unmarshal(w.Body.Bytes(), &renewedTokenResponse); err != nil {
		t.Fatalf("Failed to unmarshal renewed token response: %v", err)
	}
	renewedAccessToken, ok := renewedTokenResponse["access_token"].(string)
	if !ok {
		t.Fatal("Failed to get renewed access token from response")
	}
	renewedRefreshToken, ok := renewedTokenResponse["refresh_token"].(string)
	if !ok {
		t.Fatal("Failed to get renewed refresh token from response")
	}

	req, err = http.NewRequest(http.MethodPost, "/dashboard/api/profile", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)


	req.Header.Set("Authorization", "Bearer "+renewedAccessToken)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, refreshToken, renewedRefreshToken)
}

func TestRegisterRouter(t *testing.T) {
	err := InitServer()
	if err != nil {
		t.Fatalf("InitServer() failed: %v", err)
	}

	requestBody := `{"username":"test","password":"testtesttest"}`
	req, err := http.NewRequest(http.MethodPost, "/dashboard/api/register", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user_id")
	assert.Contains(t, w.Body.String(), "username")
	assert.Contains(t, w.Body.String(), "role")
	assert.Contains(t, w.Body.String(), "refresh_token")
	assert.Contains(t, w.Body.String(), "access_token")
}
