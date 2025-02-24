package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func ShowSignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}
