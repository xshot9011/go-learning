package main

import (
	"context"
	"fmt"
	"offersapp/routes"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {

	conn, err := connectDatabase()
	if err != nil {
		return
	}

	server := gin.Default()
	server.Use(databaseMiddleware(*conn))

	usersGroup := server.Group("users") // access url/users/[sub_route]/
	{
		usersGroup.POST("register", routes.UserRegistration)
	}

	server.Run(":8080")
}

func connectDatabase() (connection *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:g9pl6fmujiyd@localhost:5432/offersapp")
	if err != nil {
		fmt.Printf("Error connecting to db : %v", err.Error())
	}
	_ = conn.Ping(context.Background())
	return conn, err
}

func databaseMiddleware(connection pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		// attach
		c.Set("db", connection)
		c.Next()
	}
}
