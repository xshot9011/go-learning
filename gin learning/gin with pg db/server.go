package main

import (
	"log"
	"testDB/config"
	"testDB/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetRouter(router)
	config.GetConnetion()

	log.Fatal(router.Run(":8080"))
}
