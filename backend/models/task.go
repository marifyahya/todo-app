package models

import (
	"time"
)

type Task struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `json:"description"`
	Status      string     `gorm:"default:pending" json:"status"`
	Priority    string     `gorm:"default:medium" json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	UserID      uint       `json:"user_id"`
	User        User       `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt   time.Time  `json:"created_at"`
}
