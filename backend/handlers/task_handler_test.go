package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/marifyahya/todo-app/backend/database"
	"github.com/marifyahya/todo-app/backend/middleware"
	"github.com/marifyahya/todo-app/backend/models"
	"github.com/stretchr/testify/assert"
)

func generateTestToken(userID uint) string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "test-secret"
		os.Setenv("JWT_SECRET", secret)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

func TestGetTasks(t *testing.T) {
	database.SetupTestDB()
	gin.SetMode(gin.TestMode)

	// Create user
	user := models.User{Email: "task@example.com", Name: "Task User"}
	database.DB.Create(&user)

	// Create tasks for this user
	database.DB.Create(&models.Task{Title: "Task 1", UserID: user.ID, Status: "pending"})
	database.DB.Create(&models.Task{Title: "Task 2", UserID: user.ID, Status: "completed"})

	// Create task for another user
	database.DB.Create(&models.Task{Title: "Other Task", UserID: 999})

	token := generateTestToken(user.ID)

	t.Run("successful fetch tasks", func(t *testing.T) {
		r := gin.Default()
		r.GET("/tasks", middleware.AuthMiddleware(), GetTasks)

		req, _ := http.NewRequest("GET", "/tasks", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		
		data := response["data"].([]interface{})
		assert.Equal(t, 2, len(data)) // Should only get 2 tasks for this user
	})

	t.Run("filter by status", func(t *testing.T) {
		r := gin.Default()
		r.GET("/tasks", middleware.AuthMiddleware(), GetTasks)

		req, _ := http.NewRequest("GET", "/tasks?status=completed", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		data := response["data"].([]interface{})
		assert.Equal(t, 1, len(data))
		assert.Equal(t, "completed", data[0].(map[string]interface{})["status"])
	})

	t.Run("unauthorized", func(t *testing.T) {
		r := gin.Default()
		r.GET("/tasks", middleware.AuthMiddleware(), GetTasks)

		req, _ := http.NewRequest("GET", "/tasks", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
