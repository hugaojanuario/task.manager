package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string      `json"title"`
	Description string      `json"description"`
	Status      TaskManager `json"status"`
}

var Tasks []Task
