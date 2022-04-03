package main

import (
	"context"
	"fmt"
	"offersapp/routes"
	"time"

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

	//หากเราต้องการให้มี function กลางที่ต้องทำก่อนใช้ API โดยจะมีรูปแบบการใช้งานดังนี้
	// คำสั่ง r.Use จะเป็นการประกาศว่าทุกๆ route ที่อยู่ใต้บรรทัดนี้จะต้องผ่านการ execute จากบรรทัดนี้ก่อนเท่านั้น
	/*
		server.Use(exampleMiddleware())
		{
			server.POST("/test", exampleFunc)
		}
	*/

	middlewareGroup := server.Group("middleware")
	middlewareGroup.Use(exampleMiddelware())
	{
		middlewareGroup.POST("timer")
	}

	usersGroup := server.Group("users") // access url/users/[sub_route]/
	{
		usersGroup.POST("register", routes.UserRegistration)
	}

	server.Run(":8080")
}

/*
>> Middleware function จะ return gin.HandlerFunc หรือก็คือ function รูปแบบเดียวกับที่เราทำไปก่อนหน้านี้นั่นเอง
func exampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//middleware
	}
}
*/

func exampleMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input Result
		e := c.BindJSON(&input)
		if e != nil {
			fmt.Printf("error is : %v\n", e.Error())
		}
		if input.Name == "" {
			fmt.Println("Fatal. Name is empty")
			c.Abort()
		}
		c.Set("name", input.Name)
		start := time.Now()
		c.Next()
		runtime := time.Now().Sub(start).Seconds() * 1000

		c.JSON(200, gin.H{
			"status":  "success",
			"runtime": runtime,
		})
	}
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
