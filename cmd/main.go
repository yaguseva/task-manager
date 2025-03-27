package main

import (
	"taskManager/config"
	"taskManager/internal/app"
	"taskManager/pkg/logger"
)

func main() {
	config.Initialize()
	appLogger := logger.Initialize()
	if err := app.New(appLogger).Run(); err != nil {
		panic(err)
	}
}
