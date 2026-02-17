package repository

import (
	"database/sql"

	"github.com/hugaojanuario/task.manager.api/internal/domain"
)

type TaskRepository interface {
	Create(task *domain.Task) error
	FindAll() ([]domain.Task, error)
	FindByID(id string) (*domain.Task, error)
	Update(id string, task *domain.Task) error
	Delete(id string) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *domain.Task) error {
	query := `INSERT INTO tasks (title, description, status)
	          VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Status,
	).Scan(&task.ID)
}

func (r *taskRepository) FindAll() ([]domain.Task, error) {
	query := `SELECT id, title, description, status FROM tasks`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]domain.Task, 0)

	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

func (r *taskRepository) FindByID(id string) (*domain.Task, error) {
	query := `SELECT id, title, description, status FROM tasks WHERE id=$1`

	var task domain.Task
	err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) Update(id string, task *domain.Task) error {
	query := `UPDATE tasks
	          SET title=$1, description=$2, status=$3
	          WHERE id=$4`

	result, err := r.db.Exec(
		query,
		task.Title,
		task.Description,
		task.Status,
		id,
	)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *taskRepository) Delete(id string) error {
	query := `DELETE FROM tasks WHERE id=$1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
