package sqlite

import (
	"context"
	"github.com/google/uuid"
)

func (r *Repo) DeleteTask(ctx context.Context, id uuid.UUID) error {
	task, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	query := `DELETE FROM tasks WHERE id = ?`
	_, err = r.db.Exec(query, task.ID)
	return err
}
