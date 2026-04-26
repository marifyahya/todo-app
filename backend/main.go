package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marifyahya/todo-app/backend/database"
	"github.com/marifyahya/todo-app/backend/handlers"
)

func main() {
	// Initialize Database
	database.InitDB()

	// Initialize Gin Router
	r := gin.Default()

	// Routes
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
		}
	}

	log.Println("Server starting on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
