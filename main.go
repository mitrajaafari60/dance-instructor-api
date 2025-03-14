package main

import (
	"dance-instructor-api/cmd"
	"dance-instructor-api/internal/auth"
	"log"
	"os"
)



func main() {
	if os.Getenv("PRINT_TEST_TOKEN") == "true" {
		getTestToken()
		return
	}

	// Initialize JWT tools for token generation and validation
	jwtTools, err := auth.NewJWTTools("private.pem", "public.pem")
	if err != nil {
		log.Fatalf("Failed to initialize JWT tools: %v", err)
	}
	cmd.Server(jwtTools)
}

func getTestToken() {
	jwtTools, err := auth.NewJWTTools("private.pem", "public.pem")
	if err != nil {
		log.Fatal(err)
	}
	token, err := jwtTools.GenerateToken("test-user")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("JWT Token:", token)
}
