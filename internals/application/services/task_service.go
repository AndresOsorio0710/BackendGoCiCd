package services

import (
	"errors"

	"github.com/AndresOsorio0710/BackendGoCiCd/internals/core/entities"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/core/ports"
)

type ITaskService interface {
	Create(task *entities.Task) error
	GetAll() ([]*entities.Task, error)
	GetByID(id int) (*entities.Task, error)
}

type TaskService struct {
	taskRepo ports.ITaskRepository
}

func NewTaskService(taskRepo ports.ITaskRepository) ITaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) Create(task *entities.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	if task.Priority < 0 || task.Priority > 5 {
		return errors.New("priority must be between 0 and 5")
	}
	return s.taskRepo.Create(task)
}

func (s *TaskService) GetAll() ([]*entities.Task, error) {
	return s.taskRepo.GetAll()
}

func (s *TaskService) GetByID(id int) (*entities.Task, error) {
	return s.taskRepo.GetByID(id)
}
