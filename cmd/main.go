package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"taskManager/internal/app"
	"taskManager/internal/app/delivery"
	"taskManager/internal/config"
	"taskManager/internal/repository"
	"taskManager/internal/usecase"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	repo := repository.New()
	uc := usecase.New(repo)
	handlers := delivery.New(uc)
	engine := gin.Default()
	app.InitTaskEndpoints(engine, handlers)
	engine.Run(":" + config.Config.Server.Port)
}
