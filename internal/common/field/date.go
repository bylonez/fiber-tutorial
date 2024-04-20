package field

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Date time.Time

func (d Date) String() string {
	return time.Time(d).Format("2006-01-02")
}

func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date(time.Time{})
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("can't convert %T to Date", value)
	}
	*d = Date(t)
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return time.Time(d).Format("2006-01-02"), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var strDate string
	if err := json.Unmarshal(data, &strDate); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}
