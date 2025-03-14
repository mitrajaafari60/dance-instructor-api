package tests

import (
	"bytes"
	"dance-instructor-api/internal/auth" // Add JWT auth package
	"dance-instructor-api/internal/handlers"
	"dance-instructor-api/internal/middleware"
	"dance-instructor-api/internal/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndGetInstructorIntegration(t *testing.T) {
	// Initialize JWT tools
	jwtTools, err := auth.NewJWTTools("../private.pem", "../public.pem")
	if err != nil {
		t.Fatal("Failed to initialize JWT tools: ", err)
	}

	// Generate a valid JWT token
	token, err := jwtTools.GenerateToken("test-user")
	if err != nil {
		t.Fatal("Failed to generate JWT token: ", err)
	}

	// Initialize store and handler
	store := storage.NewMemoryStore()
	handler := handlers.NewInstructorHandler(store)
	router := gin.Default()

	// Use JWT middleware instead of old AuthMiddleware
	api := router.Group("/api/v1").Use(middleware.JWTAuth(jwtTools))
	api.POST("/instructors", handler.CreateInstructor)
	api.GET("/instructors/:id", handler.GetInstructorByID)

	// Create instructor
	instructor := map[string]string{
		"name":         "Integration Test",
		"bio":          "Test Bio",
		"specialty":    "Test Dance",
		"availability": "Test Times",
	}
	body, _ := json.Marshal(instructor)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/instructors", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token) // Use JWT token
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Get created instructor
	var createdInstructor struct {
		ID string `json:"id"`
	}
	json.Unmarshal(w.Body.Bytes(), &createdInstructor)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/instructors/"+createdInstructor.ID, nil)
	req.Header.Set("Authorization", "Bearer "+token) // Use JWT token
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
