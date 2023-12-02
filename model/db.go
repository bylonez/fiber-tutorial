package model

import (
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// AutoMigrate run auto migration for gorm model
	err = DBConn.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
}
