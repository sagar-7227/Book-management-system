package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	_ "modernc.org/sqlite"  // This forces use of CGO-free driver
)

var DB *gorm.DB

func ConnectDatabase() {
	d, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
