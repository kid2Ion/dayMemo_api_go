package model

import (
	"time"
)

type User struct {
	ID          string `json:"id" gorm:"praimaly_key"`
	Email       string `json:"email"`
	UserName    string `json:"user_name" validate:"min=4,max=10,alphanum"`
	DisplayName string `json:"display_name" validate:"min=1,max=10"`
	IconUrl     string `json:"icon_url"`
	Password    string `json:"password" validate:"min=6"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CreateUser(user *User) {
	db.Create(user)
}

func UpdateUser(user *User) *User {
	db.Save(user)

	return user
}

func DeleteUser(user *User) {
	db.Delete(user)
}

func FindUser(u *User) User {
	var user User
	db.Where("user_name = ?", u.UserName).First(&user)

	return user
}
