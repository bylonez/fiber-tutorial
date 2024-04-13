package common

import (
	"fmt"
)

type Error int

const (
	SystemError Error = 12300 + iota
	ParamInvalid
	CustomError
	ExportEmptyData
)

// Panic with additional msg
func (e Error) Panic(msg ...any) {
	panic(ErrorStruct{err: e, msg: msg})
}

// desc plain text description
func (e Error) desc() string {
	switch e {
	case SystemError:
		return "system error"
	case ParamInvalid:
		return "param invalid"
	case CustomError:
		return "custom %v error"
	case ExportEmptyData:
		return "export empty data"
	default:
		return "error"
	}
}

type ErrorStruct struct {
	err Error
	msg []any
}

// Error implements error interface
func (s ErrorStruct) Error() string {
	return fmt.Sprintf(s.err.desc(), s.msg...)
}
