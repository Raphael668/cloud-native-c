package models

type User struct {
	GormModel
	Email string `gorm:"type:varchar(128)"` //trons varchar(255)
}
