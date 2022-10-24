package repo

import (
	"cloud-native-c/pkg/db"
	"cloud-native-c/pkg/models"
)

type UserRepo struct {
}

func (t *UserRepo) CreateUser(user models.User) (*models.User, error) {
	ddb := db.DB
	out := &models.User{}
	*out = user

	dbc := ddb.Where("name = ?",
		out.Name).
		FirstOrCreate(&out)

	return out, dbc.Error

}
