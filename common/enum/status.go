package enum

import "fiber-tutorial/common"

type StatusEnum int

const (
	StatusAEnum StatusEnum = iota
	StatusBEnum
	StatusCEnum
)

var StatusEnums []StatusEnum

type statusStruct struct {
	name string
	desc string
}

var m map[StatusEnum]statusStruct

func (e StatusEnum) Name() string {
	return m[e].name
}

func (e StatusEnum) Desc() string {
	return m[e].desc
}

func init() {
	m = map[StatusEnum]statusStruct{
		StatusAEnum: {"a name", "a desc"},
		StatusBEnum: {"b name", "b desc"},
		StatusCEnum: {"c name", "c desc"},
	}
	StatusEnums = common.GetSortedKeys(m)
}
