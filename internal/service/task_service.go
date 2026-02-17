// @ViitoJooj
// mudei o nome para "task_service.go" para melhor semantica

package service

import (
	"github.com/hugaojanuario/task.manager.api/internal/domain"
	"github.com/hugaojanuario/task.manager.api/internal/helpers"
	"github.com/hugaojanuario/task.manager.api/internal/repository"
)

type TaskService interface {
	Create(task *domain.Task) error
	FindAll() ([]domain.Task, error)
	FindByID(id string) (*domain.Task, error)
	Update(id string, task *domain.Task) error
	Delete(id string) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) Create(task *domain.Task) error {
	if err := helpers.ValidationTask(task); err != nil {
		return err
	}

	task.Status = domain.StatusPedende
	return s.repo.Create(task)
}

func (s *taskService) FindAll() ([]domain.Task, error) {
	return s.repo.FindAll()
}

func (s *taskService) FindByID(id string) (*domain.Task, error) {
	return s.repo.FindByID(id)
}

func (s *taskService) Update(id string, task *domain.Task) error {
	if err := helpers.ValidationTask(task); err != nil {
		return err
	}
	return s.repo.Update(id, task)
}

func (s *taskService) Delete(id string) error {
	return s.repo.Delete(id)
}
