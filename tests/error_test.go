package tests

import (
	"bytes"
	"dance-instructor-api/internal/auth" 
	"dance-instructor-api/internal/handlers"
	"dance-instructor-api/internal/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorHandling(t *testing.T) {
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

	store := storage.NewMemoryStore()
	handler := handlers.NewInstructorHandler(store)
	router := SetupTestRouter(handler)

	tests := []struct {
		name           string
		method         string
		url            string
		body           interface{}
		expectedStatus int
	}{
		{"Invalid JSON", "POST", "/api/v1/instructors", "{invalid}", http.StatusBadRequest},
		{"Missing fields", "POST", "/api/v1/instructors", map[string]string{"name": "only"}, http.StatusBadRequest},
		{"Non-existent ID", "GET", "/api/v1/instructors/nonexistent", nil, http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var body []byte
			if tt.body != nil {
				body, _ = json.Marshal(tt.body)
			}
			req, _ := http.NewRequest(tt.method, tt.url, bytes.NewBuffer(body))
			req.Header.Set("Authorization", "Bearer "+token) // Use JWT token
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
