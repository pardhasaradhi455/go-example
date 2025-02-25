package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test default endpoint
func TestDefaultEndpoint(t *testing.T) {
	// Initialize Gin router
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "the default end point")
	})

	// Create a test HTTP request
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert status code and response body
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "\"the default end point\"", w.Body.String())
}

// Test /message endpoint
func TestMessageEndpoint(t *testing.T) {
	// Initialize Gin router
	router := gin.Default()
	code = 42 // setting a fixed value for code for consistency in the test
	router.GET("/message", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "your application is deployed successfully",
			"code":    code,
		})
	})

	// Create a test HTTP request
	req, _ := http.NewRequest("GET", "/message", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert response body and check if code is correct
	expectedBody := "{\"code\":42,\"message\":\"your application is deployed successfully\"}"
	assert.Equal(t, expectedBody, w.Body.String())
}
