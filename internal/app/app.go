package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"taskManager/internal/app/delivery"
	"taskManager/internal/repository"
	"taskManager/internal/usecase"
)

type ILogger interface {
	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
}

type App struct {
	engine *gin.Engine
	logger ILogger
}

func New(logger ILogger) *App {
	engine := gin.Default()
	handlers := delivery.New(usecase.New(repository.New()))
	InitTaskEndpoints(engine, handlers)
	return &App{
		engine: engine,
		logger: logger,
	}
}

func (a *App) Run() error {
	return a.engine.Run(":" + viper.GetString("server.port"))
}
