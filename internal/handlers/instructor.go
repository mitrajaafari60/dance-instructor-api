package handlers

import (
	"dance-instructor-api/internal/models"
	"dance-instructor-api/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// InstructorHandler manages HTTP requests for instructor operations.
type InstructorHandler struct {
	store *storage.MemoryStore
}

// NewInstructorHandler creates a new InstructorHandler with the given storage.
func NewInstructorHandler(store *storage.MemoryStore) *InstructorHandler {
	return &InstructorHandler{store: store}
}

// GetAllInstructors retrieves a list of all instructors.
// @Summary Get all instructors
// @Description Retrieves a list of all instructors
// @Tags instructors
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Instructor "List of instructors"
// @Router /instructors [get]
func (h *InstructorHandler) GetAllInstructors(c *gin.Context) {
	instructors := h.store.GetAll()
	c.JSON(http.StatusOK, instructors)
}

// GetInstructorByID retrieves a specific instructor by their ID.
// @Summary Get an instructor by ID
// @Description Retrieves a specific instructor by their ID
// @Tags instructors
// @Accept  json
// @Produce  json
// @Param id path string true "Instructor ID"
// @Success 200 {object} models.Instructor "Instructor details"
// @Failure 404 {object} map[string]string "Instructor not found"
// @Router /instructors/{id} [get]
func (h *InstructorHandler) GetInstructorByID(c *gin.Context) {
	id := c.Param("id")
	instructor, exists := h.store.Get(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		return
	}
	c.JSON(http.StatusOK, instructor)
}

// CreateInstructor creates a new instructor record.
// @Summary Create a new instructor
// @Description Creates a new instructor record
// @Tags instructors
// @Accept  json
// @Produce  json
// @Param instructor body models.Instructor true "Instructor object"
// @Success 201 {object} models.Instructor "Created instructor"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Router /instructors [post]
func (h *InstructorHandler) CreateInstructor(c *gin.Context) {
	var instructor models.Instructor
	if err := c.ShouldBindJSON(&instructor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	instructor.ID = uuid.New().String()
	h.store.Create(instructor.ID, instructor)
	c.JSON(http.StatusCreated, instructor)
}

// UpdateInstructor updates an existing instructor record.
// @Summary Update an instructor
// @Description Updates an existing instructor record
// @Tags instructors
// @Accept  json
// @Produce  json
// @Param id path string true "Instructor ID"
// @Param instructor body models.Instructor true "Instructor object"
// @Success 200 {object} models.Instructor "Updated instructor"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 404 {object} map[string]string "Instructor not found"
// @Router /instructors/{id} [put]
func (h *InstructorHandler) UpdateInstructor(c *gin.Context) {
	id := c.Param("id")
	var instructor models.Instructor
	if err := c.ShouldBindJSON(&instructor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	instructor.ID = id
	if !h.store.Update(id, instructor) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		return
	}
	c.JSON(http.StatusOK, instructor)
}

// DeleteInstructor removes an instructor by their ID.
// @Summary Delete an instructor
// @Description Removes an instructor by their ID
// @Tags instructors
// @Accept  json
// @Produce  json
// @Param id path string true "Instructor ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string "Instructor not found"
// @Router /instructors/{id} [delete]
func (h *InstructorHandler) DeleteInstructor(c *gin.Context) {
	id := c.Param("id")
	if !h.store.Delete(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
