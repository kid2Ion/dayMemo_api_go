package model

import "time"

type Memory struct {
	ID int `json:"id" gorm:"praimaly_key"`
	UserId string `json:"user_id"`
	content string `json:"content"`
	ImageUrl string `json:"image_url"`
	CreatedAt time.Time
}

type MemoryList []Memory

func Create