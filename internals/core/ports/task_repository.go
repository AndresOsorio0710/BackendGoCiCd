package ports

import "github.com/AndresOsorio0710/BackendGoCiCd/internals/core/entities"

type ITaskRepository interface {
	Create(task *entities.Task) error
	GetAll() ([]*entities.Task, error)
	GetByID(id int) (*entities.Task, error)
}
