package impl

import (
	"fiber-tutorial/internal/database"
	"fiber-tutorial/internal/pkg/field"
	"fiber-tutorial/internal/service/userservice"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint       `gorm:"primarykey"`
	Name     string     `gorm:"not null;size:32"`
	Birthday field.Date `gorm:"type:date;not null"`
	Gender   string     `gorm:"not null;size:32"`
}

func init() {
	// AutoMigrate run auto migration for gorm database
	err := database.DBConn.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
}

func ListUser(query *userservice.UserQuery) []*User {
	var users []*User
	database.DBConn.Order("created_at desc").Limit(query.PageSize).Offset(query.Offset()).Find(&users)
	return users
}

func CreateUser(user *User) *User {
	database.DBConn.Create(&user)
	return user
}

func UpdateUser(user *User) *User {
	database.DBConn.Updates(&user)
	return user
}
