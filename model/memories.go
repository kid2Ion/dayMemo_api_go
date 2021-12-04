package model

import (
	"fmt"
	"time"
)

type Memory struct {
	UID       int    `json:"uid"`
	ID        int    `json:"id" gorm:"praimaly_key"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ImageUrl  string `json:"image_url"`
	CreatedAt time.Time
}

type MemoryList []Memory

func CreateMemory(memory *Memory) {
	db.Create(memory)
}

func FindMemories(m *Memory) MemoryList {
	var memoryList MemoryList
	db.Where(m).Find(&memoryList)

	return memoryList
}

func UpdateMemory(m *Memory) error {
	rows := db.Model(m).Update(map[string]interface{}{
		"content":  m.Content,
		"imageurl": m.ImageUrl,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("could not find memory: %v", m)
	}

	return nil
}

func DeleteMemory(m *Memory) error {
	if rows := db.Where(m).Delete(&Memory{}).RowsAffected; rows == 0 {
		return fmt.Errorf("could not find memory:%v to delete", m)
	}

	return nil
}
