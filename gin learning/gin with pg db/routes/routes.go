package routes

import (
	"net/http"

	controllers "testDB/controller"

	"github.com/gin-gonic/gin"
)

// SetRouter set the path for routing
func SetRouter(router *gin.Engine) {
	router.GET("", welcome)
	router.GET("/todos", controllers.GetAllTodos)
	router.POST("/todo", controllers.CreateTodo)
	router.GET("/todo/:todoId", controllers.GetSingleTodo)
	router.PUT("/todo/:todoId", controllers.EditTodo)
	router.DELETE("/todo/:todoId", controllers.DeleteTodo)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to golang and gin",
	})
}
