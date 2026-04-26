package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marifyahya/todo-app/backend/database"
	"github.com/marifyahya/todo-app/backend/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestRegister(t *testing.T) {
	database.SetupTestDB()
	gin.SetMode(gin.TestMode)

	t.Run("successful registration", func(t *testing.T) {
		r := gin.Default()
		r.POST("/register", Register)

		reqBody := map[string]string{
			"email":    "test@example.com",
			"password": "password123",
			"name":     "Test User",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "User registered successfully", response["message"])
		assert.Equal(t, "success", response["status"])
		
		data := response["data"].(map[string]interface{})
		assert.Equal(t, "test@example.com", data["email"])
		assert.Equal(t, "Test User", data["name"])
	})

	t.Run("invalid email", func(t *testing.T) {
		r := gin.Default()
		r.POST("/register", Register)

		reqBody := map[string]string{
			"email":    "invalid-email",
			"password": "password123",
			"name":     "Test User",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("missing fields", func(t *testing.T) {
		r := gin.Default()
		r.POST("/register", Register)

		reqBody := map[string]string{
			"email": "test@example.com",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestLogin(t *testing.T) {
	database.SetupTestDB()
	gin.SetMode(gin.TestMode)

	// Create a user first
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.User{
		Email:    "login@example.com",
		Password: string(hashedPassword),
		Name:     "Login User",
	}
	database.DB.Create(&user)

	t.Run("successful login", func(t *testing.T) {
		r := gin.Default()
		r.POST("/login", Login)

		reqBody := map[string]string{
			"email":    "login@example.com",
			"password": "password123",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		
		data := response["data"].(map[string]interface{})
		assert.NotEmpty(t, data["token"])
	})

	t.Run("wrong password", func(t *testing.T) {
		r := gin.Default()
		r.POST("/login", Login)

		reqBody := map[string]string{
			"email":    "login@example.com",
			"password": "wrongpassword",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
