package main

import (
	"cloud-native-c/config"
	"cloud-native-c/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

const confPath = "./config.json"

func main() {

	cfg, err := config.New(confPath)
	if err != nil {
		log.Fatal(err)
	}

	addr := ":" + cfg.Server.Port
	log.Printf("======= Server start to listen (%s) and serve =======\n", addr)
	r := Router()
	r.Run(addr)
}

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/hello", hello)

	// .. //
	user := router.Group("/user") //本人
	user.POST("/deposit", controllers.Deposit)
	user.POST("/withdrawal", controllers.Withdraw)

	users := router.Group("/users")
	users.POST("/:id/deposit", controllers.DepositTo)
	users.POST("/:id/withdrawal", controllers.WithdrawFrom)

	return router
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{ // response json
		"message": "hello world",
	})
}
