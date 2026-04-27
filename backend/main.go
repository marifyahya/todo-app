package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/marifyahya/todo-app/backend/database"
	"github.com/marifyahya/todo-app/backend/handlers"
	"github.com/marifyahya/todo-app/backend/middleware"
	"github.com/marifyahya/todo-app/backend/utils"
)

func main() {
	// Define migration flag
	migrateFlag := flag.String("migrate", "", "Run migrations: 'up' or 'down'")
	flag.Parse()

	// If migrate flag is provided, run migration and exit
	if *migrateFlag != "" {
		database.RunMigrations(*migrateFlag)
		os.Exit(0)
	}

	// Normal startup: Initialize Database and run 'up' migration by default
	database.InitDB("up")

	// Initialize Gin Router
	r := gin.Default()

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to API Todo Application",
		})
	})

	// 404 Custom Response
	r.NoRoute(func(c *gin.Context) {
		utils.ErrorResponse(c, http.StatusNotFound, "Route not found")
	})

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		tasks := api.Group("/tasks")
		tasks.Use(middleware.AuthMiddleware())
		{
			tasks.GET("", handlers.GetTasks)
			tasks.POST("", handlers.CreateTask)
			tasks.PUT("/:id", handlers.UpdateTask)
		}
	}

	log.Println("Server starting on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
