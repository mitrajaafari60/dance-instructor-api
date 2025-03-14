package cmd

import (
	_ "dance-instructor-api/docs"
	"dance-instructor-api/internal/auth"
	"dance-instructor-api/internal/config"
	"dance-instructor-api/internal/handlers"
	"dance-instructor-api/internal/middleware"
	"dance-instructor-api/internal/storage"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Dance Instructor API
// @version 1.0
// @description API for managing dance instructors
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@danceinstructor.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Server(jwtTools *auth.JWTTools) {
	// Initialize in-memory storage for instructors
	store := storage.NewMemoryStore()

	// Set Gin mode based on environment (default to debug if not set)
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode // Default to debug for development
	}
	gin.SetMode(ginMode)

	// Create Gin router with custom configuration
	r := gin.New()
	r.Use(gin.Logger())   // Add logging middleware
	r.Use(gin.Recovery()) // Add recovery middleware for panic handling

	// Configure trusted proxies (set to none for safety; adjust as needed)
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// Initialize instructor handler with storage
	instructorHandler := handlers.NewInstructorHandler(store)

	// Define protected API routes under /api/v1 with JWT authentication
	api := r.Group("/api/v1").Use(middleware.JWTAuth(jwtTools))
	{
		api.GET("/instructors", instructorHandler.GetAllInstructors)       // List all instructors
		api.GET("/instructors/:id", instructorHandler.GetInstructorByID)   // Get instructor by ID
		api.POST("/instructors", instructorHandler.CreateInstructor)       // Create new instructor
		api.PUT("/instructors/:id", instructorHandler.UpdateInstructor)    // Update existing instructor
		api.DELETE("/instructors/:id", instructorHandler.DeleteInstructor) // Delete instructor
	}

	// Add Swagger UI route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server on the configured port
	log.Printf("Starting server on %s in %s mode", config.GetPort(), ginMode)
	if err := r.Run(config.GetPort()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
