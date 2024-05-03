package ex

type Exception int

const (
	SystemError Exception = 12300 + iota
	ParamInvalid
	ExportEmptyData
)

var m = map[Exception]string{
	SystemError:     "system exception",
	ParamInvalid:    "param invalid",
	ExportEmptyData: "export empty data",
}

// Panic with additional msg
func (e Exception) Panic(msg ...any) {
	panic(ExceptionPanic{Ex: e, Msg: msg})
}

// Desc plain text description
func (e Exception) Desc() string {
	return m[e]
}

func AddEx(ex Exception, desc string) {
	m[ex] = desc
}
