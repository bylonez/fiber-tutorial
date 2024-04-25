package enum

import (
	"fiber-tutorial/pkg/map"
)

type StatusEnum int

const (
	StatusAEnum StatusEnum = iota
	StatusBEnum
	StatusCEnum
)

type statusStruct struct {
	name string
	desc string
}

var m = map[StatusEnum]statusStruct{
	StatusAEnum: {"a name", "a desc"},
	StatusBEnum: {"b name", "b desc"},
	StatusCEnum: {"c name", "c desc"},
}

var StatusEnums = _map.GetSortedKeys(m)

func (e StatusEnum) Name() string {
	return m[e].name
}

func (e StatusEnum) Desc() string {
	return m[e].desc
}
