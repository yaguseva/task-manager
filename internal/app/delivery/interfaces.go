package delivery

import (
	"context"
	"github.com/google/uuid"
	"taskManager/internal/entity"
)

type IUseCase interface {
	CreateTask(ctx context.Context, task entity.Task) (uuid.UUID, error)
	GetFilteredTasks(ctx context.Context, status *bool, priority *int) (map[uuid.UUID]entity.Task, error)
	UpdateTask(ctx context.Context, id uuid.UUID, task entity.Task) (entity.Task, error)
	DeleteTask(ctx context.Context, id uuid.UUID) error
}
