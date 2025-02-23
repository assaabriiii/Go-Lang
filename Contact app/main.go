package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var contacts = []Contact{
	{ID: 1, Name: "Amir Sabri", Email: "john@example.com", Phone: "123-456-7890"},
}

func main() {
	router := gin.Default()

	router.GET("/contacts", func(c *gin.Context) {
		c.JSON(http.StatusOK, contacts)
	})

	router.GET("/contacts/:id", func(c *gin.Context) {
		IdStr := c.Param("id")
		id, err := strconv.Atoi(IdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID"})
			return
		}

		for _, contact := range contacts {
			if contact.ID == id {
				c.JSON(http.StatusOK, contact)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"Error": "Contact not found"})
	})

	router.POST("/contacts", func(c *gin.Context) {
		var newContact Contact
		err := c.ShouldBindJSON(&newContact)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to bind JSON"})
			return
		}

		newContact.ID = len(contacts) + 1
		contacts = append(contacts, newContact)

		c.JSON(http.StatusOK, newContact)
	})

	router.PUT("/contacts/:id", func(c *gin.Context) {
		IdStr := c.Param("id")
		id, err := strconv.Atoi(IdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID"})
			return
		}

		var updatedContact Contact
		err = c.ShouldBindJSON(&updatedContact)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to bind JSON"})
			return
		}

		for i, contact := range contacts {
			if contact.ID == id {
				contacts[i].Name = updatedContact.Name
				contacts[i].Email = updatedContact.Email
				contacts[i].Phone = updatedContact.Phone
				c.JSON(http.StatusOK, updatedContact)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"Error": "Contact not found"})
	})

	router.DELETE("/contacts/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID"})
			return
		}

		for i, contact := range contacts {
			if contact.ID == id {
				contacts = append(contacts[:i], contacts[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"Message": "Contact successfully deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"Error": "Contact not found"})
	})

	router.Run(":8080")
}
