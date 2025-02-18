package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// todo app
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{
	{ID: 1, Title: "Learn Go language", Done: false},
	{ID: 2, Title: "Go to your classes", Done: true},
}

func main() {
	// Create gin router
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Route to serve the HTML page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	router.POST("/add", func(c *gin.Context) {
		title := c.PostForm("title")

		if title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Title is required"})
			return
		}

		newTodo := Todo{
			ID:    len(todos) + 1,
			Title: title,
			Done:  false,
		}
		todos = append(todos, newTodo)

		c.Redirect(http.StatusFound, "/")
	})

	router.POST("/done/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			println("Error converting id to string")
			panic(err)
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Done = true
				break
			}
		}
		c.Redirect(http.StatusFound, "/")
	})

	router.Run(":8080")
}
