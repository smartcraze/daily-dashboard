package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status" gorm:"default:pending"`
	DueDate     string `json:"due_date"`
}

type DailyBrief struct {
	gorm.Model
	Weather string `json:"weather"`
	News    string `json:"news"`
	Quote   string `json:"quote"`
}
