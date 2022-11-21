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

func Bigger10Add(a, b int) (int, error) {
	if a <= 10 || b <= 10 {
		return 0, fmt.Errorf("invald")
	}

	return a + b, nil
}
