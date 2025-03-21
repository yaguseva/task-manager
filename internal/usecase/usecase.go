package usecase

import (
	"github.com/google/uuid"
	"taskManager/internal/entity"
)

type UseCase struct {
	repo IDatabase
}

func New(repo IDatabase) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) CreateTask(task entity.Task) (uuid.UUID, error) {
	task.ID = uuid.New()
	return task.ID, uc.repo.CreateTask(task)
}

func (uc *UseCase) GetFilteredTasks(status *bool, priority *int) (map[uuid.UUID]entity.Task, error) {
	tasks, err := uc.repo.GetFilteredTasks(status, priority)
	if err != nil {
		return nil, err
	}
	result := make(map[uuid.UUID]entity.Task, len(tasks))
	for _, task := range tasks {
		result[task.ID] = task
	}
	return result, nil
}

func (uc *UseCase) UpdateTask(id uuid.UUID, task entity.Task) (entity.Task, error) {
	return uc.repo.UpdateTask(id, task)
}

func (uc *UseCase) DeleteTask(id uuid.UUID) error {
	return uc.repo.DeleteTask(id)
}
