package routes

import (
	"net/http"
	"regisapp/controllers"

	"github.com/gin-gonic/gin"
)

// SetRouter set the path for routing
func SetRouter(router *gin.Engine) {
	router.GET("", welcome)
	router.POST("register", controllers.Registration)
	// router.POST("login", controllers.Login)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to vulnerability web application",
	})
}
