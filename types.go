package gatsby

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// NullInt32 represents an int64 that may be null.
// NullInt32 implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullInt32 struct {
	Int32 int32
	Valid bool // Valid is true if Int32 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullInt32) Scan(value interface{}) error {
	if value == nil {
		n.Int32, n.Valid = 0, false
		return nil
	}
	n.Valid = true

	if val, ok := value.(int32); ok {
		n.Int32 = val
		return nil
	}
	var s = asString(value)
	val, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return fmt.Errorf("converting string %q to a int32: %v", s, err)
	}
	n.Int32 = int32(val)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullInt32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int32, nil
}

func asString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	return fmt.Sprintf("%v", src)
}
