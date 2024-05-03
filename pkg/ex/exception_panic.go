package ex

import (
	"fmt"
)

type ExceptionPanic struct {
	Ex  Exception
	Msg []any
}

// Error implements ex interface
func (s ExceptionPanic) Error() string {
	return fmt.Sprintf(s.Ex.Desc(), s.Msg...)
}
