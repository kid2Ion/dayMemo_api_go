package model

import (
	"time"
)

type Users struct {
	ID          string `json:"id" gorm:"praimaly_key"`
	Email       string `json:"email"`
	UserName    string `json:"user_name" validate:"min=4,max=10,alphanum"`
	DisplayName string `json:"display_name" validate:"min=1,max=10"`
	IconBase64  string `json:"icon_base64"`
	IconUrl     string
	Password    string `json:"password" validate:"min=6"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CreateUser(user *Users) {
	db.Create(user)
}

func UpdateUser(user *Users) *Users {
	db.Save(user)

	return user
}

func DeleteUser(user *Users) {
	db.Delete(user)
}

func FindUser(u *Users) Users {
	var user Users
	db.Where("user_name = ?", u.UserName).First(&user)

	return user
}

func FindUserByUid(u *Users) Users {
	var user Users
	db.Where("id = ?", u.ID).First(&user)

	return user
}
