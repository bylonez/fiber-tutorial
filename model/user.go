package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string    `gorm:"not null;size:32"`
	Birthday time.Time `gorm:"type:date;not null"`
	Gender   string    `gorm:"not null;size:32"`
}

func List() *[]*User {
	var users []*User
	DBConn.Find(&users)
	return &users
}

func Create(user *User) *User {
	DBConn.Create(&user)
	return user
}
