package database

import (
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	DBConn *gorm.DB
)

func init() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

}
