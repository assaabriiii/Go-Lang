package controllers

import (
	"login/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users []models.User

func Signup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	newUser := models.User{
		Username: username,
		Password: password,
	}

	users = append(users, newUser)

	c.HTML(http.StatusOK, "login.html", gin.H{
		"message": "Signup successful! Please login.",
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	for _, user := range users {
		if user.Username == username && user.Password == password {
			c.HTML(http.StatusOK, "welcome.html", gin.H{"username": username})
			return
		}
	}

	c.HTML(http.StatusUnauthorized, "login.html", gin.H{
		"error": "Invalid credentials. Please try again.",
	})
}
