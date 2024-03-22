package user

import "fiber-tutorial/common/field"

type userExportVO struct {
	Id       uint       `excel:"head:i"`
	Name     string     `excel:"head:n"`
	Birthday field.Date `excel:"head:bir"`
	Gender   string     `excel:"head:gen"`
}
