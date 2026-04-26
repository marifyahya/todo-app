package database

import (
	"github.com/marifyahya/todo-app/backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupTestDB initializes an in-memory SQLite database for testing purposes.
// It also runs migrations for all models.
func SetupTestDB() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}

	// Migrate all models here
	db.AutoMigrate(&models.User{}, &models.Task{})

	DB = db
}
