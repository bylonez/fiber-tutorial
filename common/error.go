package common

import (
	"fmt"
	"strings"
)

type Error int

const (
	SystemError Error = 12300 + iota
	ParamInvalid
	CustomError
)

// Panic with additional msg
func (e Error) Panic(msg ...string) {
	panic(ErrorStruct{err: e, msg: msg})
}

// desc plain text description
func (e Error) desc() string {
	switch e {
	case SystemError:
		return "system error"
	case ParamInvalid:
		return "param invalid"
	default:
		return "error"
	}
}

// fmtDesc format text description
func (e Error) fmtDesc() string {
	switch e {
	case CustomError:
		return "custom %v error"
	default:
		return ""
	}
}

type ErrorStruct struct {
	err Error
	msg []string
}

// Error implements error interface
func (s ErrorStruct) Error() string {
	fmtDesc := s.err.fmtDesc()
	desc := s.err.desc()
	if len(s.msg) == 0 {
		// mo additional msg
		if desc != "" {
			return desc
		} else {
			return fmtDesc
		}
	} else {
		if fmtDesc != "" {
			// format desc
			x := make([]any, len(s.msg))
			for i, v := range s.msg {
				x[i] = v
			}
			return fmt.Sprintf(fmtDesc, x...)
		} else {
			// desc + s.msg
			return desc + ", " + strings.Join(s.msg, ", ")
		}
	}
}
