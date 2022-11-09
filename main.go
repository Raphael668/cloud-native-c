package main

import (
	"cloud-native-c/pkg/app/server"
	"cloud-native-c/pkg/config"
	"fmt"
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

func PositiveAdd(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("invald")
	}

	return a + b, nil
}
