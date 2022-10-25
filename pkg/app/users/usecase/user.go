package usecase

import (
	"cloud-native-c/pkg/app/users/repo"
	"cloud-native-c/pkg/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var userRepo repo.User

type User struct {
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (*User) Register(user *models.User) (*models.User, error) {

	pw, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = pw
	return userRepo.CreateUser(user)
}

func (*User) Login(name, password string) (*models.User, error) {

	user, err := userRepo.GetUser(name)
	if err != nil {
		return nil, fmt.Errorf("can't not find user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {

		return nil, fmt.Errorf("invalid password error")
	}

	return user, nil
}
