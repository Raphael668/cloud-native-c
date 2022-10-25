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

	"golang.org/x/crypto/bcrypt"
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

	err = seed()
	if err != nil {
		log.Fatal(err)
	}
}

func migration(db *gorm.DB) {

	log.Println("Run DB Migration --> Start")

	db.AutoMigrate(&models.User{})

	log.Println("Run DB Migration --> Done")

}

func seed() error {

	log.Println("Create Seed --> Start")

	user, err := loadSeed(seedPath)
	if err != nil {
		return err
	}

	IO := repo.User{}
	IO.CreateUser(*user) //char(32) will be error

	fmt.Printf("%+v", user)

	log.Println("Create Seed --> Done")
	return nil
}

func loadSeed(path string) (*models.User, error) {
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

	pw, err := HashPassword(seed.User.Password)
	if err != nil {
		return nil, err
	}
	seed.User.Password = pw
	return &seed.User, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
