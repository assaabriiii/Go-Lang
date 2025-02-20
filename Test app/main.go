package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	test := ""

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"test": test,
		})
	})

	router.GET("/getter", func(c *gin.Context) {
		test = c.Query("test")
		test = "I'm not here for plot " + test

		c.HTML(http.StatusOK, "get.html", gin.H{
			"test": test,
		})
	})

	// Post needs to be called from another page and processed in other page
	router.POST("/poster", func(c *gin.Context) {
		fmt.Println("POST /post route hit") // Debugging line
		test = c.PostForm("test")
		if test == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Test parameter is required"})
			return
		}

		test = "I'm not here for plot " + test

		c.HTML(http.StatusOK, "post.html", gin.H{
			"test": test,
		})
	})

	// router.DELETE("/delete", func(c *gin.Context) {
	// 	test = ""
	// 	c.HTML(http.StatusOK, "delete.html", gin.H{
	// 		"test": test,
	// 	})
	// })

	router.Run(":8080")
}
