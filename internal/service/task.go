package service

import (
	"github.com/hugaojanuario/task.manager.api/model"
	"gopkg.in/validator.v2"
)

func ValidationTask(t *model.Task) error {
	if err := validator.Validate(&t); err != nil {
		return err
	}
	return nil
}
