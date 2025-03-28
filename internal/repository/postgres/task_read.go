package postgres

import (
	"context"
	"github.com/google/uuid"
	"strconv"
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
	row := r.db.QueryRowContext(ctx, `SELECT * FROM tasks WHERE id = $1`, id)
	var task entity.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority)
	return task, err
}

func (r *Repo) getFilteredTasksQuery(status *bool, priority *int) (string, []any) {
	query := `SELECT * from tasks`
	conditions := []string{}
	args := []any{}

	argNum := 0
	if status != nil {
		argNum++
		conditions = append(conditions, `status = $`+strconv.Itoa(argNum))
		args = append(args, *status)
	}
	if priority != nil {
		argNum++
		conditions = append(conditions, `priority = $`+strconv.Itoa(argNum))
		args = append(args, *priority)
	}
	if status != nil || priority != nil {
		query += ` WHERE ` + strings.Join(conditions, " AND ")
	}
	return query, args
}
