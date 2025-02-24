package routes

import (
	"mod-app/controllers"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.GET("/products", controllers.GetProducts)
	router.POST("/users", controllers.AddUser)
	router.POST("/products", controllers.AddProduct)
	router.GET("/users/:id", controllers.GetUser)
	router.GET("/products/:id", controllers.GetProduct)
	router.GET("/users/template", controllers.ShowUsersPage)
}
