package models

type User struct {
	GormModel
	Name string `gorm:"unique;type:varchar(128)"` //trons varchar(255)
}
