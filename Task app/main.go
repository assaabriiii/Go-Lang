package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var tasks = []Task{
	{ID: 1, Title: "Learn Go", Desc: "Complete the Go tutorial"},
	{ID: 2, Title: "Build a REST API", Desc: "Create a task manager API with Gin"},
}

func main() {
	router := gin.Default()

	router.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, tasks)
	})

	router.GET("/tasks/:id", func(c *gin.Context) {
		IdStr := c.Param("id")
		id, err := strconv.Atoi(IdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Converting ID to string failed"})
			return
		}

		for _, task := range tasks {
			if task.ID == id {
				c.JSON(http.StatusOK, task)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
	})

	router.POST("/tasks", func(c *gin.Context) {
		var newTask Task
		err := c.ShouldBindJSON(&newTask)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "ShouldBindJSON failed"})
			return
		}

		newTask.ID = len(tasks) + 1
		tasks = append(tasks, newTask)

		c.JSON(http.StatusCreated, newTask)
	})

	router.PUT("/tasks/:id", func(c *gin.Context) {
		IdStr := c.Param("id")
		id, err := strconv.Atoi(IdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to convert"})
			return
		}

		var updatedNewTask Task
		err = c.ShouldBindJSON(&updatedNewTask)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request body"})
			return
		}

		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Title = updatedNewTask.Title
				tasks[i].Desc = updatedNewTask.Desc
				c.JSON(http.StatusOK, tasks[i])
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
	})

	router.DELETE("/tasks/:id", func(c *gin.Context) {
		IdStr := c.Param("id")
		id, err := strconv.Atoi(IdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to convert"})
			return
		}

		for i, task := range tasks {
			if task.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"Message": "Task successfully deleted"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
	})

	router.Run(":8080")
}
