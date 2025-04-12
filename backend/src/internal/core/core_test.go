package core

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInitServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := InitServer()
	if err != nil {
		t.Fatalf("InitServer() failed: %v", err)
	}
}

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := InitServer()
	if err != nil {
		t.Fatalf("InitServer() failed: %v", err)
	}

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
}
