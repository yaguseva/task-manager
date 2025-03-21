package usecase

import (
	"github.com/google/uuid"
	"taskManager/internal/entity"
)

type IDatabase interface {
	CreateTask(task entity.Task) error
	GetFilteredTasks(status *bool, priority *int) ([]entity.Task, error)
	UpdateTask(id uuid.UUID, task entity.Task) (entity.Task, error)
	DeleteTask(id uuid.UUID) error
}
