package server

import (
	"cloud-native-c/pkg/config"
	"cloud-native-c/pkg/db"
	"log"
)

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
