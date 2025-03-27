package sqlite

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
	query := `UPDATE tasks set title = ?, description = ?, status = ?, priority = ? WHERE id = ?`
	_, err = r.db.ExecContext(ctx, query, task.Title, task.Description, task.Status, task.Priority, fromDb.ID)
	if err != nil {
		return result, err
	}
	return r.Get(ctx, id)
}
