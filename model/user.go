package model

import (
	"fiber-tutorial/common/field"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint       `gorm:"primarykey"`
	Name     string     `gorm:"not null;size:32"`
	Birthday field.Date `gorm:"type:date;not null"`
	Gender   string     `gorm:"not null;size:32"`
}

func ListUser() *[]*User {
	var users []*User
	DBConn.Find(&users)
	return &users
}

func CreateUser(user *User) *User {
	DBConn.Create(&user)
	return user
}

func UpdateUser(user *User) *User {
	DBConn.Updates(&user)
	return user
}
