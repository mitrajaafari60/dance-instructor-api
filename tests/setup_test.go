package tests

import (
	"dance-instructor-api/internal/auth"
	"dance-instructor-api/internal/handlers"
	"dance-instructor-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupTestRouter(handler *handlers.InstructorHandler) *gin.Engine {
	// Initialize JWT tools
	jwtTools, err := auth.NewJWTTools("../private.pem", "../public.pem")
	if err != nil {
		panic("Failed to initialize JWT tools: " + err.Error())
	}

	r := gin.Default()
	api := r.Group("/api/v1").Use(middleware.JWTAuth(jwtTools))
	{
		api.GET("/instructors", handler.GetAllInstructors)
		api.GET("/instructors/:id", handler.GetInstructorByID)
		api.POST("/instructors", handler.CreateInstructor)
		api.PUT("/instructors/:id", handler.UpdateInstructor)
		api.DELETE("/instructors/:id", handler.DeleteInstructor)
	}
	return r
}
