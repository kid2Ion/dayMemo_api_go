package model

import "time"

type User struct {
	ID          string `json:"id" gorm:"praimaly_key"`
	MailAddress string `json:"mail_address"`
	Name        string `json:"name"`
	IconUrl     string `json:"icon_url"`
	Password    string `json:"password"`
	CreatedAt   time.Time
}

func CreateUser(user *User) {
	db.Create(user)
}

func FindUser(u *User) User {
	var user User
	db.Where(u).First(&user)

	return user
}
