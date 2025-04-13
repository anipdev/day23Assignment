package main

import (
	"day23Assignment/config"
	"day23Assignment/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	config.InitDB()

	// Set up routes
	router := routes.SetupRouter()

	// Run the server
	router.Run(":8080")
}
