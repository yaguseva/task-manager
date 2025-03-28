package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	_ "taskManager/migrations"
)

type Repo struct {
	db *sql.DB
}

func New() *Repo {
	user := viper.GetString("db.postgres.user")
	password := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("db.postgres.name")
	sslMode := viper.GetString("db.postgres.sslmode")
	connStr := "user=" + user + " password=" + password + " dbname=" + dbName + " sslmode=" + sslMode + " host=localhost port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = UpMigrations(db); err != nil {
		panic(err)
	}

	return &Repo{db: db}
}

func UpMigrations(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(db, "./migrations"); err != nil {
		return err
	}
	return nil
}
