package main

import (
	"log"
	"regisapp/models"
	"regisapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.SetRouter(server)
	models.GetConnectionDB()

	log.Fatal(server.Run(":8080"))
}
