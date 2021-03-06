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
	err := c.ShouldBindJSON(&user)  // จะเป็นการ bind ข้อมูลมาเก็บไว้ใน structure โดยใช้คำสั่ง BindJSON
	//นั้นจะสามารถทำได้ครั้งเดียว นั่นหมายความว่า หาก bind ไปแล้วข้อมูลส่วนนั้นใน request จะหายไปเลย ไม่สามารถ bind ซ้ำได้อีก
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
