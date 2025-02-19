package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var URLMap = make(map[string]string)

func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	shortCode := make([]byte, 6)
	for i := range shortCode {
		shortCode[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortCode)
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/shorten", func(c *gin.Context) {
		longURL := c.PostForm("url")
		if longURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "URL is required"})
			return
		}
		shortCode := generateShortCode()
		URLMap[shortCode] = longURL

		c.HTML(http.StatusOK, "short.html", gin.H{
			"shortURL": "http://localhost:8080/" + shortCode,
		})
	})

	router.GET("/:shortCode", func(c *gin.Context) {
		shortCode := c.Param("shortCode")
		longURL, exists := URLMap[shortCode]
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Short URL not found",
			})
			return
		}

		c.Redirect(http.StatusMovedPermanently, longURL)
	})

	router.Run(":8080")
}
