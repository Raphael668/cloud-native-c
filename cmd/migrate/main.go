package main

import (
	"cloud-native-c/pkg/app/users/repo"
	"cloud-native-c/pkg/config"
	"cloud-native-c/pkg/db"
	"cloud-native-c/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gorm.io/gorm"
)

const confPath = "./config.json"
const seedPath = "./seed.json"

func main() {

	cfg, err := config.New(confPath)
	if err != nil {
		log.Fatal(err)
	}

	d, err := db.GormOpen(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	migration(d)

	user, err := LoadSeed(seedPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", user)

	IO := repo.UserRepo{}
	IO.CreateUser(*user)
}

func migration(db *gorm.DB) {

	log.Println("Run DB Migration --> Start")

	db.AutoMigrate(&models.User{})

	log.Println("Run DB Migration --> Done")

}

func LoadSeed(path string) (*models.User, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	type Seed struct {
		User models.User `json:"User"`
	}

	var seed Seed

	err = json.Unmarshal(buf, &seed)
	if err != nil {
		return nil, err
	}

	return &seed.User, nil
}
