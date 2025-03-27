package sqlite

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	_ "modernc.org/sqlite"
	_ "taskManager/migrations"
)

type Repo struct {
	db *sql.DB
}

func New() *Repo {
	db, err := sql.Open("sqlite", viper.GetString("db.path"))
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
