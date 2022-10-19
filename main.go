package main

import (
	"cloud-native-c/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	r := Router()
	r.Run(":3000") // default localhost:8000
}

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/hello", hello)

	// .. //
	user := router.Group("/user")

	user.POST("/deposit", controllers.Deposit)
	user.POST("/withdrawal", controllers.Withdrawal)

	return router
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{ // response json
		"message": "hello world",
	})
}
