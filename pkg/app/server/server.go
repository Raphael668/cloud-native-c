package server

import (
	"cloud-native-c/pkg/config"
	"cloud-native-c/pkg/controllers"
	"cloud-native-c/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)

const confPath = "./config.json"

type Server struct {
	Config *config.Config
}

func New(cfg *config.Config) *Server {

	return &Server{
		Config: cfg,
	}

}

func (t *Server) Serve() {

	_, err := db.GormOpen(&t.Config.DB)
	if err != nil {
		log.Fatal(err)
	}
	addr := ":" + t.Config.Server.Port
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
