package impl

import (
	"fiber-tutorial/internal/database"
	"fiber-tutorial/internal/pkg/field"
	"fiber-tutorial/internal/service/userservice"
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
	database.RegModel(&User{})
}

func ListUser(query *userservice.UserQuery) []*User {
	var users []*User
	database.DB.Order("created_at desc").Limit(query.PageSize).Offset(query.Offset()).Find(&users)
	return users
}

func CreateUser(user *User) *User {
	database.DB.Create(&user)
	return user
}

func UpdateUser(user *User) *User {
	database.DB.Updates(&user)
	return user
}
