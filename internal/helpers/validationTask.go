package helpers

import (
	"errors"

	"github.com/hugaojanuario/task.manager.api/internal/domain"
)

func ValidationTask(task *domain.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}

	if len(task.Title) < 3 {
		return errors.New("title must have at least 3 characters")
	}

	if task.Description == "" {
		return errors.New("description is required")
	}

	return nil
}
