package common

import "fmt"

type Error int

const (
	SYSTEM_ERROR Error = 12300 + iota
)

func (e Error) Error() string {
	switch e {
	case SYSTEM_ERROR:
		return "SYSTEM_ERROR"
	default:
		panic(fmt.Sprintf("unknown Error, %v", e))
	}
}
