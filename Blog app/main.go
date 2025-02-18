package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts = []Post{
	{ID: 1, Title: "First Post", Body: "This is the first blog post."},
	{ID: 2, Title: "Second Post", Body: "This is the second blog post."},
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"posts": posts,
		})
	})

	router.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add.html", nil)
	})

	router.POST("/add", func(c *gin.Context) {
		title := c.PostForm("title")
		body := c.PostForm("body")

		if title == "" || body == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Title and body must be provided"})
			return
		}

		newPost := Post{
			ID:    len(posts) + 1,
			Title: title,
			Body:  body,
		}
		posts = append(posts, newPost)

		c.Redirect(http.StatusFound, "/")
	})

	router.GET("/post/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid id"})
			return
		}

		var post Post
		for _, p := range posts {
			if p.ID == id {
				post = p
				break
			}
		}

		if post.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Post not found"})
			return
		}

		c.HTML(http.StatusOK, "post.html", gin.H{
			"post": post,
		})

	})

	router.Run(":8080")
}
