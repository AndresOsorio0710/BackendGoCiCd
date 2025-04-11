package services_test

import (
	"context"
	"testing"

	"github.com/AndresOsorio0710/BackendGoCiCd/internals/application/services"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/config"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/core/entities"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/infrastructure/dbcontext"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) (*repository.TaskRepository, *dbcontext.DbContext) {
	cfg := config.PostgresConfig{
		Host:     "localhost",
		Port:     5432, // Puerto de testing
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	}

	ctx, err := dbcontext.NewDbContext(cfg)
	_, err = ctx.OpenConnection()
	assert.NoError(t, err)

	return repository.NewTaskRepository(ctx), ctx
}

func cleanDB(t *testing.T, dbCtx *dbcontext.DbContext) {
	conn, err := dbCtx.OpenConnection()
	assert.NoError(t, err)
	defer conn.Close()

	_, err = conn.ExecContext(context.Background(), `DELETE FROM tasks`)
	assert.NoError(t, err)
}

type mockRepo struct {
	storage []entities.Task
}

func (m *mockRepo) Create(task *entities.Task) error {
	id := len(m.storage) + 1
	task.ID = &id
	m.storage = append(m.storage, *task)
	return nil
}

func (m *mockRepo) GetAll() ([]entities.Task, error) {
	return m.storage, nil
}

func (m *mockRepo) GetByID(id int) (*entities.Task, error) {
	for _, t := range m.storage {
		if t.ID == &id {
			return &t, nil
		}
	}
	return nil, nil
}

func TestCreate_ValidTask(t *testing.T) {
	mock, ctx := setupTestDB(t)
	service := services.NewTaskService(mock)

	cleanDB(t, ctx)

	task := &entities.Task{
		Title:    "Test",
		Priority: 3,
	}

	err := service.Create(task)
	assert.NoError(t, err)
}

func TestCreate_InvalidPriority(t *testing.T) {
	mock, _ := setupTestDB(t)
	service := services.NewTaskService(mock)

	task := &entities.Task{
		Title:    "Test",
		Priority: 6,
	}

	err := service.Create(task)
	assert.Error(t, err)
	assert.Equal(t, "priority must be between 0 and 5", err.Error())
}

func TestCreate_EmptyTitle(t *testing.T) {
	mock, _ := setupTestDB(t)
	service := services.NewTaskService(mock)

	task := &entities.Task{
		Title:    "",
		Priority: 2,
	}

	err := service.Create(task)
	assert.Error(t, err)
	assert.Equal(t, "title is required", err.Error())
}
