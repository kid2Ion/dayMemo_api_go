package model

import (
	"strconv"
	"time"
)

type Memory struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UID       string `json:"uid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ImageUrl  string
}

type MemoryList []Memory

func CreateMemory(memory *Memory) {
	db.Create(memory)
}

func FindMemories(m *Memory, year string, month string) MemoryList {
	var memoryList MemoryList

	yearInt, _ := strconv.Atoi(year)
	monthInt, _ := (strconv.Atoi(month))
	startTime := time.Date(yearInt, time.Month(monthInt), 1, 0, 0, 0, 0, time.Local)
	endTime := startTime.AddDate(0, 1, 0)

	db.Where(m).Where("created_at BETWEEN ? AND ?", startTime, endTime).Find(&memoryList)

	return memoryList
}

func FindMemory(m *Memory) Memory {
	var memory Memory
	db.Where(m).First(&memory)

	return memory
}

func UpdateMemory(memory *Memory) *Memory {
	db.Save(memory)

	return memory
}

func DeleteMemory(memory *Memory) {
	db.Delete(memory)
}
