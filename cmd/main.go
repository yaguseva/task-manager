package main

import (
	"taskManager/internal/app"
	"taskManager/internal/config"
	"taskManager/pkg/logger"
)

func main() {
	err := config.Initialize()
	if err != nil {
		panic(err)
	}
	appLogger := logger.Initialize()
	if err := app.New(appLogger).Run(); err != nil {
		panic(err)
	}
}
