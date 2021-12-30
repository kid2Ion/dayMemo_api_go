package model

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Memory struct {
	gorm.Model
	UID      string `json:"uid"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"`
}

type MemoryList []Memory

func CreateMemory(memory *Memory) {
	db.Create(memory)
}

func FindMemories(m *Memory, year_month string) MemoryList {
	var memoryListAll MemoryList
	db.Where(m).Find(&memoryListAll)

	thisYearMonth := year_month
	var nextYearMonth string
	var lastYearMonth string

	// 12月と1月は前後の月でyearを変更するロジックが必要
	rep := regexp.MustCompile(`\s*-\s*`)
	result := rep.Split(year_month, -1)

	thisYearInt, _ := strconv.Atoi(result[0])
	nextYear := strconv.Itoa(thisYearInt + 1)
	lastYear := strconv.Itoa(thisYearInt - 1)
	thisMonthInt, _ := strconv.Atoi(result[1])
	nextMonth := strconv.Itoa(thisMonthInt + 1)
	if len(nextMonth) == 1 {
		nextMonth = "0" + nextMonth
	}
	lastMonth := strconv.Itoa(thisMonthInt - 1)
	if len(lastMonth) == 1 {
		lastMonth = "0" + lastMonth
	}

	if result[1] == "12" {
		nextYearMonth = nextYear + "-01"
		lastYearMonth = result[0] + "-11"
	} else if result[1] == "01" {
		nextYearMonth = result[0] + "-02"
		lastYearMonth = lastYear + "-12"
	} else {
		nextYearMonth = result[0] + "-" + nextMonth
		lastYearMonth = result[0] + "-" + lastMonth
	}

	var memoryList MemoryList
	for _, v := range memoryListAll {
		MemorycreatedAt := v.CreatedAt.Format("2006-01")
		if MemorycreatedAt == thisYearMonth || MemorycreatedAt == nextYearMonth || MemorycreatedAt == lastYearMonth {
			memoryList = append(memoryList, v)
		}
	}
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
