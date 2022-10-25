package repo

import (
	"cloud-native-c/pkg/db"
	"cloud-native-c/pkg/models"
)

type User struct {
}

func (t *User) CreateUser(user models.User) (*models.User, error) {
	ddb := db.DB
	out := &models.User{}
	*out = user

	dbc := ddb.Where("name = ?",
		out.Name).
		FirstOrCreate(&out)

	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return out, nil

}

func (t *User) GetUser(name string) (*models.User, error) {
	ddb := db.DB
	out := &models.User{}

	dbc := ddb.Where("name = ?",
		name).
		First(&out)

	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return out, nil

}
