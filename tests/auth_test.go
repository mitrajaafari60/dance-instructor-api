package tests

import (
	"dance-instructor-api/internal/auth"
	"dance-instructor-api/internal/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestJWTAuthMiddleware tests the JWT authentication middleware.
func TestJWTAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Initialize JWT tools
	jwtTools, err := auth.NewJWTTools("../private.pem", "../public.pem")
	if err != nil {
		t.Fatal("Failed to initialize JWT tools: ", err)
	}

	// Set up test router
	router := gin.Default()
	router.GET("/test", middleware.JWTAuth(jwtTools), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Generate a valid token
	validToken, err := jwtTools.GenerateToken("test-user")
	if err != nil {
		t.Fatal("Failed to generate token: ", err)
	}

	// Define test cases
	tests := []struct {
		name           string
		token          string
		expectedStatus int
	}{
		{"Valid token", "Bearer " + validToken, http.StatusOK},
		{"Invalid token", "Bearer invalid-token", http.StatusUnauthorized},
		{"No token", "", http.StatusUnauthorized},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", tt.token)
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
