package repository

import (
	"context"

	"github.com/AndresOsorio0710/BackendGoCiCd/internals/core/entities"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/infrastructure/dbcontext"
)

type TaskRepository struct {
	dbCtx *dbcontext.DbContext
}

func NewTaskRepository(ctx *dbcontext.DbContext) *TaskRepository {
	return &TaskRepository{dbCtx: ctx}
}

func (r *TaskRepository) Create(task *entities.Task) error {
	conn, err := r.dbCtx.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := `INSERT INTO tasks (title, description, priority) VALUES ($1, $2, $3)`
	_, err = conn.ExecContext(context.Background(), query, task.Title, task.Description, task.Priority)
	return err
}

func (r *TaskRepository) GetAll() ([]*entities.Task, error) {
	conn, err := r.dbCtx.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `SELECT id, title, description, priority, created_at FROM tasks ORDER BY created_at DESC`
	rows, err := conn.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entities.Task
	for rows.Next() {
		var t entities.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Priority, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}

	return tasks, nil
}

func (r *TaskRepository) GetByID(id int) (*entities.Task, error) {
	conn, err := r.dbCtx.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `SELECT id, title, description, priority, created_at FROM tasks WHERE id = $1`
	row := conn.QueryRowContext(context.Background(), query, id)

	var t entities.Task
	err = row.Scan(&t.ID, &t.Title, &t.Description, &t.Priority, &t.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
