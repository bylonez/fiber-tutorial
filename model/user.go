package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   int    `gorm:"primaryKey"`
	Name string `json:"name"`
}

func List() *[]User {
	var users []User
	DBConn.Find(&users)
	return &users
}

func Create(user *User) {
	DBConn.Create(&user)
}
