package fb_visual

import (
	"database/sql/driver"
)

func (receiver Type) Value() (driver.Value, error) {
	return int64(receiver.value)
}
