package handlers

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/marifyahya/todo-app/backend/database"
	"github.com/marifyahya/todo-app/backend/models"
	"github.com/marifyahya/todo-app/backend/utils"
)

type CreateTaskRequest struct {
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	Priority    *string    `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
}

type UpdateTaskRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Priority    *string    `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
}

func GetTasks(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	var tasks []models.Task
	query := database.DB.Where("user_id = ?", userID)

	// Filter by status
	status := c.Query("status")
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Filter by priority
	priority := c.Query("priority")
	if priority != "" {
		query = query.Where("priority = ?", priority)
	}

	// Sorting
	sort := c.Query("sort")
	switch sort {
	case "created_at_asc":
		query = query.Order("created_at asc")
	case "created_at_desc":
		query = query.Order("created_at desc")
	case "due_date_asc":
		query = query.Order("due_date asc")
	case "due_date_desc":
		query = query.Order("due_date desc")
	default:
		query = query.Order("created_at desc")
	}

	if err := query.Find(&tasks).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch tasks")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Tasks retrieved successfully", tasks)
}

func CreateTask(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Validation failed", gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		UserID:      userID.(uint),
		Status:      "pending",
	}

	if err := database.DB.Create(&task).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create task")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Task created successfully", task)
}

func UpdateTask(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	id := c.Param("id")
	var task models.Task

	// Find task and check ownership
	if err := database.DB.First(&task, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Task not found")
		return
	}

	if task.UserID != userID.(uint) {
		utils.ErrorResponse(c, http.StatusForbidden, "You do not have permission to update this task")
		return
	}

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Validation failed", gin.H{"error": err.Error()})
		return
	}

	// Update fields
	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.Priority != nil {
		updates["priority"] = req.Priority
	}
	if req.DueDate != nil {
		updates["due_date"] = req.DueDate
	}

	if err := database.DB.Model(&task).Updates(updates).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update task")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Task updated successfully", task)
}

func DeleteTask(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	id := c.Param("id")
	var task models.Task

	// Find task and check ownership
	if err := database.DB.First(&task, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Task not found")
		return
	}

	if task.UserID != userID.(uint) {
		utils.ErrorResponse(c, http.StatusForbidden, "You do not have permission to delete this task")
		return
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Task deleted successfully", nil)
}
