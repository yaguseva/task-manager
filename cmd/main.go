package main

import (
	"taskManager/config"
	"taskManager/internal/app"
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
