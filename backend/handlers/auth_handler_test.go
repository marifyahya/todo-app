package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marifyahya/todo-app/backend/database"
	"github.com/stretchr/testify/assert"
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

		userData := response["user"].(map[string]interface{})
		assert.Equal(t, "test@example.com", userData["email"])
		assert.Equal(t, "Test User", userData["name"])
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
