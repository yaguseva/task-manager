package sqlite

import (
	"context"
	"taskManager/internal/entity"
)

func (r *Repo) CreateTask(ctx context.Context, task entity.Task) error {
	query := `INSERT INTO tasks (ID, TITLE, DESCRIPTION, STATUS, PRIORITY) VALUES (?, ?, ?, ?, ?);`

	_, err := r.db.ExecContext(ctx, query, task.ID, task.Title, task.Description, task.Status, task.Priority)
	if err != nil {
		return err
	}

	return nil
}
