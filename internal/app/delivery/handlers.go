package delivery

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"taskManager/internal/entity"
	"time"
)

type Handler struct {
	uc IUseCase
}

func New(usecase IUseCase) *Handler {
	return &Handler{uc: usecase}
}

func (h *Handler) CreateTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task entity.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.uc.CreateTask(ctx, task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) GetFilteredTasks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	statusParam := c.Query("status")
	priorityParam := c.Query("priority")
	c.Header("Cache-Control", "public, max-age=3600")

	status, err := h.parseStatusParam(statusParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	priority, err := h.parsePriorityParam(priorityParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := h.uc.GetFilteredTasks(ctx, status, priority)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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
	res, err := h.uc.UpdateTask(ctx, id, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) DeleteTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.uc.DeleteTask(ctx, id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) parseStatusParam(statusParam string) (*bool, error) {
	var status *bool
	if statusParam != "" {
		status = new(bool)
		var err error
		*status, err = strconv.ParseBool(statusParam)
		if err != nil {
			return nil, err
		}
	}
	return status, nil
}

func (h *Handler) parsePriorityParam(priorityParam string) (*int, error) {
	var priority *int
	if priorityParam != "" {
		priority = new(int)
		var err error
		*priority, err = strconv.Atoi(priorityParam)
		if err != nil {
			return nil, err
		}
	}
	return priority, nil
}
