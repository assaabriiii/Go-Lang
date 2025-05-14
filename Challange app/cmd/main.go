package main

import (
	"cmd/internal/handlers"
	"cmd/internal/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	taskRepo := repository.NewTaskRepository()
	taskHandlers := handlers.NewTaskHandlers(taskRepo)
	engine := gin.Default()
	engine.Use(Logger())

	engine.GET("/tasks", taskHandlers.GetAllTasks)
	engine.GET("/tasks/:id", taskHandlers.GetTaskByID)
	engine.POST("/tasks", taskHandlers.CreateTask)
	engine.PUT("/tasks/:id", taskHandlers.UpdateTask)
	engine.DELETE("/tasks/:id", taskHandlers.DeleteTask)

	engine.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("%s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}
