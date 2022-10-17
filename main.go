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

	// .. //
	user := router.Group("/user")

	user.POST("/deposit", controllers.Deposit)
	user.POST("/withdraw", controllers.Withdraw)

	return router
}
