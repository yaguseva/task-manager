package delivery

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"strconv"
	"taskManager/internal/config"
	"taskManager/internal/entity"
	"taskManager/internal/repository"
)

type Handler struct {
	uc IUseCase
}

func New(usecase IUseCase) *Handler {
	return &Handler{uc: usecase}
}

func (h *Handler) SaveDataToFile(c *gin.Context) {
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

	for _, task := range repository.Storage {
		if err := encoder.Encode(task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, "data saved successfully")
}

func (h *Handler) CreateTask(c *gin.Context) {
	var task entity.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.uc.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) GetFilteredTasks(c *gin.Context) {
	statusParam := c.Query("status")
	priorityParam := c.Query("priority")
	c.Header("Cache-Control", "public, max-age=3600")

	var err error
	var status *bool
	if statusParam != "" {
		status = new(bool)
		*status, err = strconv.ParseBool(statusParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	var priority *int
	if priorityParam != "" {
		priority = new(int)
		*priority, err = strconv.Atoi(priorityParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	result, err := h.uc.GetFilteredTasks(status, priority)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task entity.Task
	if err = c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.uc.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) DeleteTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.uc.DeleteTask(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
