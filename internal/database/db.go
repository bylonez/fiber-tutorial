package database

import (
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	DB     *gorm.DB
	models []any
)

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("sqlite3.db"))
	if err != nil {
		log.Fatal("failed to connect database")
	}
	migrationModel()
}

func RegModel(model any) {
	models = append(models, model)
}

func migrationModel() {
	for _, model := range models {
		err := DB.AutoMigrate(model)
		if err != nil {
			log.Fatal(err)
		}
	}
}
