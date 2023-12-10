package field

import (
	"database/sql/driver"
	"time"
)

type DateTime string

// 实现 sql.Scanner 接口，Scan 将 value 扫描至
func (j *DateTime) Scan(value interface{}) error {
	*j = DateTime((value.(time.Time)).Format("2006-01-02 15:04:05"))
	return nil
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j DateTime) Value() (driver.Value, error) {
	return string(j), nil
}
