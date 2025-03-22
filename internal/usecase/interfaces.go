package usecase

import (
	"context"
	"github.com/google/uuid"
	"taskManager/internal/entity"
)

type IDatabase interface {
	CreateTask(ctx context.Context, task entity.Task) error
	GetFilteredTasks(ctx context.Context, status *bool, priority *int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, id uuid.UUID, task entity.Task) (entity.Task, error)
	DeleteTask(ctx context.Context, id uuid.UUID) error
}
