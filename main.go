package main

import (
	"cloud-native-c/app/server"
	"cloud-native-c/config"
	"log"
)

const confPath = "./config.json"

func main() {

	cfg, err := config.New(confPath)
	if err != nil {
		log.Fatal(err)
	}

	svr := server.New(cfg)
	svr.Serve()

}
