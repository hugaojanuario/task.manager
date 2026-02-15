package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string      `json:"title" validate:"min=4,max=30,regexp=^[a-zA-Z ]+$"`
	Description string      `json:"description" validate:"min=10,max=40,regexp=^[a-zA-Z ]+$"`
	Status      TaskManager `json:"status"`
}

var Tasks []Task
