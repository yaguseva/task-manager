package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTasksTable, downCreateTasksTable)
}

func upCreateTasksTable(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
	id TEXT PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT,
	status BOOLEAN,
	priority INTEGER);`

	_, err := tx.Exec(query)
	if err != nil {
		panic(err)
	}
	return nil
}

func downCreateTasksTable(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE tasks;`

	_, err := tx.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
