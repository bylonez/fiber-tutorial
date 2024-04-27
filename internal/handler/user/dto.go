package user

import "fiber-tutorial/pkg/field"

type userExportVO struct {
	Id       uint       `excel:"head:id"`
	Name     string     `excel:"head:na"`
	Birthday field.Date `excel:"head:bi"`
	Gender   string     `excel:"head:ge"`
}

type userImportCO struct {
	Id   uint   `excel:"head:id"`
	Name string `excel:"head:na"`
	//Birthday field.Date `excel:"head:bi"`
	Gender string `excel:"head:ge"`
}
