package delivery

import (
	"github.com/google/uuid"
	"taskManager/internal/entity"
)

type IUseCase interface {
	CreateTask(task entity.Task) (uuid.UUID, error)
	GetFilteredTasks(status *bool, priority *int) (map[uuid.UUID]entity.Task, error)
	UpdateTask(id uuid.UUID, task entity.Task) (entity.Task, error)
	DeleteTask(id uuid.UUID) error
}
