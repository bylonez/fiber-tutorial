package common

import "fmt"

type Error int

const (
	SystemError Error = 12300 + iota
)

func (e Error) Error() string {
	switch e {
	case SystemError:
		return "SYSTEM_ERROR"
	default:
		panic(fmt.Sprintf("unknown Error, %v", e))
	}
}
