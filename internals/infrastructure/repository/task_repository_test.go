package repository_test

import (
	"testing"

	"github.com/AndresOsorio0710/BackendGoCiCd/internals/config"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/core/entities"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/infrastructure/dbcontext"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *repository.TaskRepository {
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

	return repository.NewTaskRepository(ctx)
}

func TestCreateAndGetAllTasks(t *testing.T) {
	repo := setupTestDB(t)

	task := &entities.Task{
		Title:       "Prueba",
		Description: "Descripción de prueba",
		Priority:    3,
	}

	err := repo.Create(task)
	assert.NoError(t, err)

	tasks, err := repo.GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, tasks)

	last := tasks[0]
	assert.Equal(t, task.Title, last.Title)
	assert.Equal(t, task.Description, last.Description)
	assert.Equal(t, task.Priority, last.Priority)
}

func TestGetByID(t *testing.T) {
	repo := setupTestDB(t)

	task := &entities.Task{
		Title:       "Por ID",
		Description: "Buscar por ID",
		Priority:    2,
	}

	err := repo.Create(task)
	assert.NoError(t, err)

	tasks, _ := repo.GetAll()
	last := tasks[0]

	found, err := repo.GetByID(*last.ID)
	assert.NoError(t, err)
	assert.Equal(t, last.Title, found.Title)
}
