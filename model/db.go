package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open("sqlite3", "db/sample.db")
	if err != nil {
		panic("failed to connect datebase")
	}
	// defaultでdb名が複数形になるのを阻止
	db.SingularTable(true)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Memory{})
}
