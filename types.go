package gatsby

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// NullInt32 represents an int32 that may be null.
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

// NullInt16 represents an int64 that may be null.
// NullInt16 implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullInt16 struct {
	Int16 int16
	Valid bool // Valid is true if Int16 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullInt16) Scan(value interface{}) error {
	if value == nil {
		n.Int16, n.Valid = 0, false
		return nil
	}
	n.Valid = true

	if val, ok := value.(int16); ok {
		n.Int16 = val
		return nil
	}
	var s = asString(value)
	val, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return fmt.Errorf("converting string %q to a int16: %v", s, err)
	}
	n.Int16 = int16(val)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullInt16) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int16, nil
}

// NullInt8 represents an int64 that may be null.
// NullInt8 implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullInt8 struct {
	Int8  int8
	Valid bool // Valid is true if Int8 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullInt8) Scan(value interface{}) error {
	if value == nil {
		n.Int8, n.Valid = 0, false
		return nil
	}
	n.Valid = true

	if val, ok := value.(int8); ok {
		n.Int8 = val
		return nil
	}
	var s = asString(value)
	val, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return fmt.Errorf("converting string %q to a int8: %v", s, err)
	}
	n.Int8 = int8(val)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullInt8) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int8, nil
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
