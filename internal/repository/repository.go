package repository

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
	"taskManager/config"
	_ "taskManager/migrations"
)

type Repo struct {
	db *sql.DB
}

func New() *Repo {
	db, err := sql.Open("sqlite", config.Config.DBPath)
	if err != nil {
		panic(err)
	}

	if err = UpMigrations(db); err != nil {
		panic(err)
	}

	return &Repo{db: db}
}

func UpMigrations(db *sql.DB) error {
	if err := goose.SetDialect("sqlite"); err != nil {
		return err
	}
	if err := goose.Up(db, "./migrations"); err != nil {
		return err
	}
	return nil
}
