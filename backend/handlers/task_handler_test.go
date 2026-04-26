package handlers

import (
	"bytes"
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

func TestCreateTask(t *testing.T) {
	database.SetupTestDB()
	gin.SetMode(gin.TestMode)

	user := models.User{Email: "create@example.com", Name: "Create User"}
	database.DB.Create(&user)
	token := generateTestToken(user.ID)

	t.Run("successful create task", func(t *testing.T) {
		r := gin.Default()
		r.POST("/tasks", middleware.AuthMiddleware(), CreateTask)

		reqBody := map[string]string{
			"title":       "New Task",
			"description": "Task Description",
			"priority":    "high",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		data := response["data"].(map[string]interface{})
		assert.Equal(t, "New Task", data["title"])
		assert.Equal(t, "high", data["priority"])
		assert.Equal(t, "pending", data["status"])
	})

	t.Run("missing title", func(t *testing.T) {
		r := gin.Default()
		r.POST("/tasks", middleware.AuthMiddleware(), CreateTask)

		reqBody := map[string]string{
			"description": "Task Description",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
