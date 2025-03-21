package app

import (
	"github.com/gin-gonic/gin"
	"taskManager/internal/app/delivery"
)

func InitTaskEndpoints(engine *gin.Engine, h *delivery.Handler) {
	engine.POST("/task", h.CreateTask)
	engine.GET("/task", h.GetFilteredTasks)
	engine.PUT("/task/:id", h.UpdateTask)
	engine.DELETE("/task/:id", h.DeleteTask)
	engine.POST("/task/save", h.SaveDataToFile)
}
