package repository

import (
	"database/sql"
	_ "modernc.org/sqlite"
	"taskManager/internal/config"
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
