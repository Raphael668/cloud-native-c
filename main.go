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

func BiggerAdd(th, n1, n2 int) (int, error) {
	if n1 <= th || n2 <= th {
		return 0, fmt.Errorf("invald")
	}

	return n1 + n2, nil
}
