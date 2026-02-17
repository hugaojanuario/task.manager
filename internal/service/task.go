package service

import (
	"github.com/hugaojanuario/task.manager.api/internal/domain"
	"gopkg.in/validator.v2"
)

func ValidationTask(t *domain.Task) error {
	if err := validator.Validate(&t); err != nil {
		return err
	}
	return nil
}
