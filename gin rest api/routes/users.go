package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegistration is made for registration of user
func UserRegistration(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user_id": "some_id",
	})
}
