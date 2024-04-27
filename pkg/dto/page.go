package dto

type PageQuery struct {
	Page     int
	PageSize int
}

func (q PageQuery) Offset() int {
	return (q.Page - 1) * q.PageSize
}
