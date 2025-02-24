package controllers

import (
	"mod-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var products = []models.Product{
	{ID: 1, Name: "Laptop", Price: 999.99},
	{ID: 2, Name: "Phone", Price: 499.99},
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func AddProduct(c *gin.Context) {
	var newProduct models.Product
	err := c.ShouldBindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
}

func GetProduct(c *gin.Context) {
	IdStr := c.Param("id")
	id, err := strconv.Atoi(IdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
}
