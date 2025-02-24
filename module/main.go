package main

import (
	"mod-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	routes.Register(router)
	router.Run(":8080")
}
