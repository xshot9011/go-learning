package main

import (
	"github.com/gin-gonic/gin"
)

// Human for example return data type
type Human struct {
	firstName string `json:"first_name"`
	lastName  string `json:"last_name"`
}

func main() {
	server := gin.Default()

	server.GET("/welcome/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"messgae": "Welcome to go server!",
		})
	})
	server.GET("/welcome2/", welcome2)
	server.POST("/create_human/", createHuman)

	server.Run(":8080")
}

func welcome2(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "New function for returning",
	})
}

func createHuman(ctx *gin.Context) {
	var human Human
	human.firstName = ctx.PostForm("first_name")
	human.lastName = ctx.PostForm("last_name")

	ctx.JSON(200, gin.H{
		"human_first_name": human.firstName,
		"human_lst_name":   human.lastName,
	})
}
