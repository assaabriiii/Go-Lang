package handlers

import (
	"cmd/internal/models"
	"cmd/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandlers struct {
	repo *repository.TaskRepository
}

func NewTaskHandlers(repo *repository.TaskRepository) *TaskHandlers {
	return &TaskHandlers{repo: repo}
}

func (h *TaskHandlers) GetAllTasks(c *gin.Context) {
	tasks, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandlers) CreateTask(c *gin.Context) {
	var req models.CreateTaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}
	if req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
	}
	id := uuid.New().String()
	task := models.Task{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}
	if err := h.repo.Create(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandlers) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.repo.GetByID(id)
	if err != nil {
		if err == repository.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandlers) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var req models.CreateTaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}
	if req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
	}

	task := models.Task{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}
	if err := h.repo.Update(task); err != nil {
		if err == repository.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandlers) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := h.repo.Delete(id)
	if err != nil {
		if err == repository.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}
	c.Status(http.StatusNoContent)
}
