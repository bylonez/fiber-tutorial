package ex

import "fiber-tutorial/pkg/ex"

const (
	CustomError ex.Exception = 20000 + iota
)

func init() {
	ex.AddEx(CustomError, "custom error")
}
