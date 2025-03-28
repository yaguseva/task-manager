package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"taskManager/internal/entity"
	"taskManager/internal/repository/postgres"
	"taskManager/internal/repository/sqlite"
)

type IDatabase interface {
	CreateTask(ctx context.Context, task entity.Task) error
	GetFilteredTasks(ctx context.Context, status *bool, priority *int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, id uuid.UUID, task entity.Task) (entity.Task, error)
	DeleteTask(ctx context.Context, id uuid.UUID) error
}

func New() IDatabase {
	dbType := viper.GetString("db.type")
	switch dbType {
	case "sqlite":
		return sqlite.New()
	case "postgres":
		return postgres.New()
	default:
		panic(fmt.Sprintf("%q: unknown db type", dbType))
	}
}
