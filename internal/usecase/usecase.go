package usecase

import (
	"context"
	"github.com/google/uuid"
	"taskManager/internal/entity"
	"taskManager/internal/repository"
)

type UseCase struct {
	repo repository.IDatabase
}

func New(repo repository.IDatabase) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) CreateTask(ctx context.Context, task entity.Task) (uuid.UUID, error) {
	task.ID = uuid.New()
	return task.ID, uc.repo.CreateTask(ctx, task)
}

func (uc *UseCase) GetFilteredTasks(ctx context.Context, status *bool, priority *int) (map[uuid.UUID]entity.Task, error) {
	tasks, err := uc.repo.GetFilteredTasks(ctx, status, priority)
	if err != nil {
		return nil, err
	}
	result := make(map[uuid.UUID]entity.Task, len(tasks))
	for _, task := range tasks {
		result[task.ID] = task
	}
	return result, nil
}

func (uc *UseCase) UpdateTask(ctx context.Context, id uuid.UUID, task entity.Task) (entity.Task, error) {
	return uc.repo.UpdateTask(ctx, id, task)
}

func (uc *UseCase) DeleteTask(ctx context.Context, id uuid.UUID) error {
	return uc.repo.DeleteTask(ctx, id)
}
