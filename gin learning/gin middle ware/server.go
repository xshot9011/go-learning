package main

import (
	"fmt"
	"time"

	"learning-middleware/api"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(mainMiddleware())

	accountGroup := server.Group("account")
	accountGroup.Use(accountMiddleware()) // using middleware only this group of router
	{
		accountGroup.POST("registration", accountRegistration)
	}

	server.Run(":8080")
}

func mainMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) { // parameter is the gin environment (context)
		fmt.Println("this is inside main middelware")
	}
}

func accountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("this is inside account middleware")
		var regInfo api.RegistrationInfomation
		err := c.BindJSON(&regInfo) // extract JSON from request obj and keep them in regInfo
		if err != nil {
			fmt.Printf("error: %v\n", err.Error())
			c.JSON(400, gin.H{
				"status": "fail",
				"error":  err.Error(),
			})
			return
		}

		if regInfo.Username == "" {
			fmt.Println("username field is required")
			c.Abort() // คือสั่งให้ process จบ ณ ตรงนั้น
		}

		c.Set("new_user", regInfo)
		// จะเป็นการ set key และ value เข้าไปใน context
		// เพื่อทำให้สามารถนำค่านั้นไปใช้ที่อื่นที่ c ไปถึงได้ (เช่น middleware กับ API function)

		startTimer := time.Now()
		c.Next()
		// คือการสั่งให้ระบบข้ามไปทำใน function ต่อไป (อาจจะเป็น middleware อีกตัวหรือส่วน API เลยก็ได้)
		// แล้วจึงกลับมาทำส่วน middleware ต่อหลังจาก API process เสร็จสิ้น
		runTime := time.Now().Sub(startTimer).Seconds() * 1000

		c.JSON(200, gin.H{
			"msg":        "successfully craete account",
			"runtime":    runTime,
			"infomation": regInfo,
		})
	}
}

func accountRegistration(c *gin.Context) {
	fmt.Println("this is inside account reaccountRegistrationgistration function")
	newUser, exist := c.Get("new_user")
	//  เป็นการ get ข้อมูลที่ได้มีการ set ไว้ก่อนแล้วด้วย key โดยที่ function จะ
	// return value ที่ set ไว้พร้อมกับ boolean 1 ตัวที่บอกว่า key นั้น exist หรือเปล่า
	if !exist {
		c.JSON(500, gin.H{
			"msg": "Cannot get new_user infomation",
		})
		c.Abort()
	}

	fmt.Printf("new_user's type : %T\n", newUser)
	fmt.Println(newUser)
}
