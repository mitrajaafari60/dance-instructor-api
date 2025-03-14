package tests

import (
	"bytes"
	"dance-instructor-api/internal/handlers"
	"dance-instructor-api/internal/models"
	"dance-instructor-api/internal/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test handler function directly
func TestCreateInstructorHandler(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := handlers.NewInstructorHandler(store)

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	instructor := models.Instructor{
		Name:         "Test Dancer",
		Bio:          "Test Bio",
		Specialty:    "Test Specialty",
		Availability: "Test Availability",
	}
	body, _ := json.Marshal(instructor)
	c.Request, _ = http.NewRequest("POST", "/instructors", bytes.NewBuffer(body))

	handler.CreateInstructor(c)

	assert.Equal(t, http.StatusCreated, c.Writer.Status())
}

// Test storage layer
func TestMemoryStoreCreate(t *testing.T) {
	store := storage.NewMemoryStore()
	instructor := models.Instructor{
		ID:           "1",
		Name:         "Test",
		Bio:          "Bio",
		Specialty:    "Specialty",
		Availability: "Availability",
	}

	store.Create(instructor.ID, instructor)
	result, exists := store.Get(instructor.ID)

	assert.True(t, exists)
	assert.Equal(t, instructor, result)
}
