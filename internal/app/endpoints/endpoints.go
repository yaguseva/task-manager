package endpoints

import (
	"github.com/gin-gonic/gin"
	"taskManager/internal/app/handlers"
)

func InitTaskEndpoints(engine *gin.Engine) {
	engine.POST("/task", handlers.CreateTask)
	engine.GET("/task", handlers.GetFilteredTasks)
	engine.PUT("/task/:id", handlers.UpdateTask)
	engine.DELETE("/task/:id", handlers.DeleteTask)
	engine.POST("/task/save", handlers.SaveDataToFile)
}
