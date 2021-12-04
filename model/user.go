package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID          int    `json:"id" gorm:"praimaly_key"`
	Name        string `json:"name"`
	IconUrl     string `json:"icon_url"`
	MailAddress string `json:"mail_address"`
	Password    string `json:"password"`
}

func CreateUser(user *User) {
	db.Create(user)
}

func FindUser(u *User) User {
	var user User
	db.Where(u).First(&user)

	return user
}
