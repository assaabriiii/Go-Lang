package routes

import (
	"login/controllers"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.Signup)
	router.GET("/login", controllers.ShowLoginPage)
	router.GET("/signup", controllers.ShowSignupPage)
}
