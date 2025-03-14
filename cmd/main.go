package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"taskManager/internal/app/config"
	"taskManager/internal/app/db"
	"taskManager/internal/app/endpoints"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.LoadData()
	if err != nil {
		log.Fatal(err)
		return
	}
	engine := gin.Default()
	endpoints.InitTaskEndpoints(engine)
	engine.Run(":" + config.Config.Server.Port)
}
