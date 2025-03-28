package postgres

import (
	"context"
	"github.com/google/uuid"
	"taskManager/internal/entity"
)

func (r *Repo) UpdateTask(ctx context.Context, id uuid.UUID, task entity.Task) (entity.Task, error) {
	var result entity.Task
	fromDb, err := r.Get(ctx, id)
	if err != nil {
		return result, err
	}
	query := `UPDATE tasks set title = $1, description = $2, status = $3, priority = $4 WHERE id = $5`
	_, err = r.db.ExecContext(ctx, query, task.Title, task.Description, task.Status, task.Priority, fromDb.ID)
	if err != nil {
		return result, err
	}
	return r.Get(ctx, id)
}
