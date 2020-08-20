package main

import (
	"github.com/gin-gonic/gin"
)

// Human for example return data type
type Human struct {
	FirstName string `json:"first_name"` // exported field
	LastName  string `json:"last_name"`  // exported field
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
	human.FirstName = ctx.PostForm("first_name")
	human.LastName = ctx.PostForm("last_name")

	ctx.JSON(200, gin.H{
		// if the field is not exported, you will have to do this manually
		"human_first_name": human.FirstName,
		"human_lst_name":   human.LastName,

		// if the field is exported, JSON tag will work for you
		"human": human,
	})
}
