package controllers

import (
	"mod-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
		Products: []models.Product{
			{ID: 1, Name: "Laptop", Price: 999.99, UserID: 1},
			{ID: 2, Name: "Phone", Price: 499.99, UserID: 1},
		},
	},
	{
		ID:    2,
		Name:  "Bob",
		Email: "bob@example.com",
		Products: []models.Product{
			{ID: 3, Name: "Tablet", Price: 299.99, UserID: 2},
			{ID: 4, Name: "Headphones", Price: 199.99, UserID: 2},
		},
	},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func AddUser(c *gin.Context) {
	var newUser models.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)
}

func GetUser(c *gin.Context) {
	IdStr := c.Param("id")
	id, err := strconv.Atoi(IdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func ShowUsersPage(c *gin.Context) {
	c.HTML(http.StatusOK, "users.html", users)
}
