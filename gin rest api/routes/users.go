package routes

import (
	"fmt"
	"net/http"
	"offersapp/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

// UserRegistration is made for registration of user
func UserRegistration(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	err = user.Register(&conn)
	if err != nil {
		fmt.Println("Error in Register()")
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	token, err := user.GetAuthToken()
	if err == nil{
		c.JSON(http.StatusOK, gin.H{
			"token": token
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": "some_id",
	})
}
