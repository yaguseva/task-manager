package repository

import (
	"context"
	"github.com/google/uuid"
	"strings"
	"taskManager/internal/entity"
)

func (r *Repo) GetFilteredTasks(ctx context.Context, status *bool, priority *int) ([]entity.Task, error) {
	query, args := r.getFilteredTasksQuery(status, priority)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *Repo) Get(ctx context.Context, id uuid.UUID) (entity.Task, error) {
	row := r.db.QueryRowContext(ctx, `SELECT * FROM tasks WHERE id = ?`, id)
	var task entity.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority)
	return task, err
}

func (r *Repo) getFilteredTasksQuery(status *bool, priority *int) (string, []any) {
	query := `SELECT * from tasks`
	conditions := []string{}
	args := []any{}

	if status != nil {
		conditions = append(conditions, `status = ?`)
		args = append(args, *status)
	}
	if priority != nil {
		conditions = append(conditions, `priority = ?`)
		args = append(args, *priority)
	}
	if status != nil || priority != nil {
		query += ` WHERE ` + strings.Join(conditions, " AND ")
	}
	return query, args
}
