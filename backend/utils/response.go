package utils

import (
	"github.com/gin-gonic/gin"
)

// SuccessResponse handles uniform success responses
func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

// ErrorResponse handles uniform error responses (401, 404, 500, etc)
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
	})
}

// ValidationErrorResponse handles validation errors (400)
func ValidationErrorResponse(c *gin.Context, message string, errors interface{}) {
	c.JSON(400, gin.H{
		"status":  "error",
		"message": message,
		"errors":  errors,
	})
}
