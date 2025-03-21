package repository

import (
	"database/sql"
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
	"strings"
	"taskManager/internal/config"
	"taskManager/internal/entity"
)

type Repo struct {
	db *sql.DB
}

func New() *Repo {
	db, err := sql.Open("sqlite", config.Config.DBPath)
	if err != nil {
		panic(err)
	}

	initQuery := `CREATE TABLE IF NOT EXISTS tasks (
	id TEXT PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT,
	status BOOLEAN,
	priority INTEGER);`

	_, err = db.Exec(initQuery)
	if err != nil {
		panic(err)
	}

	return &Repo{db: db}
}

func (r *Repo) CreateTask(task entity.Task) error {
	query := `INSERT INTO tasks (ID, TITLE, DESCRIPTION, STATUS, PRIORITY) VALUES (?, ?, ?, ?, ?);`

	_, err := r.db.Exec(query, task.ID, task.Title, task.Description, task.Status, task.Priority)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetFilteredTasks(status *bool, priority *int) ([]entity.Task, error) {
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
	rows, err := r.db.Query(query, args...)
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

func (r *Repo) Get(id uuid.UUID) (entity.Task, error) {
	row := r.db.QueryRow(`SELECT * FROM tasks WHERE id = ?`, id)
	var task entity.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority)
	return task, err
}

func (r *Repo) UpdateTask(id uuid.UUID, task entity.Task) (entity.Task, error) {
	var result entity.Task
	fromDb, err := r.Get(id)
	if err != nil {
		return result, err
	}
	query := `UPDATE tasks set title = ?, description = ?, status = ?, priority = ? WHERE id = ?`
	_, err = r.db.Exec(query, task.Title, task.Description, task.Status, task.Priority, fromDb.ID)
	if err != nil {
		return result, err
	}
	return r.Get(id)
}

func (r *Repo) DeleteTask(id uuid.UUID) error {
	task, err := r.Get(id)
	if err != nil {
		return err
	}
	query := `DELETE FROM tasks WHERE id = ?`
	_, err = r.db.Exec(query, task.ID)
	return err
}
