package models

// Instructor represents a dance instructor in the system.
type Instructor struct {
	ID           string `json:"id"`                              // Unique identifier for the instructor
	Name         string `json:"name" binding:"required"`         // Instructor's full name
	Bio          string `json:"bio" binding:"required"`          // Brief biography of the instructor
	Specialty    string `json:"specialty" binding:"required"`    // Dance specialty (e.g., Ballet, Hip-Hop)
	Availability string `json:"availability" binding:"required"` // Available teaching times
}
