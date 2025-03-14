package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"strconv"
	"taskManager/internal/app/config"
	"taskManager/internal/app/db"
	"taskManager/internal/app/models"
)

func SaveDataToFile(c *gin.Context) {
	file, err := os.Create(config.Config.FileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	for _, task := range db.Storage {
		if err := encoder.Encode(task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, "data saved successfully")
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = uuid.New()
	db.Storage[task.ID] = task
	c.JSON(http.StatusCreated, task.ID)
}

func GetFilteredTasks(c *gin.Context) {
	statusParam := c.Query("status")
	priorityParam := c.Query("priority")
	c.Header("Cache-Control", "public, max-age=3600")

	if statusParam == "" && priorityParam == "" {
		c.JSON(http.StatusOK, db.Storage)
		return
	}

	var err error
	var status bool
	if statusParam != "" {
		status, err = strconv.ParseBool(statusParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	var priority int
	if priorityParam != "" {
		priority, err = strconv.Atoi(priorityParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	var result = make(map[uuid.UUID]models.Task)
	for _, task := range db.Storage {
		if (statusParam != "" && status != task.Status) || (priorityParam != "" && priority != task.Priority) {
			continue
		} else {
			result[task.ID] = task
		}
	}
	c.JSON(http.StatusOK, result)
}

func UpdateTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task models.Task
	if err = c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, exists := db.Storage[id]
	if !exists {
		c.Status(http.StatusNotFound)
		return
	}
	task.ID = id
	db.Storage[task.ID] = task
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, exists := db.Storage[id]
	if !exists {
		c.Status(http.StatusNotFound)
		return
	}
	delete(db.Storage, id)
	c.Status(http.StatusOK)
}
