package nullable

import (
	"database/sql"
	"strconv"
)

// NullInt64 wraps sql.NullInt64 to allow it to be serializable to/from XML
// via TextMarshaler and TextUnmarshaler
type NullInt64 struct {
	sql.NullInt64
}

// NewNullInt64 creats a new NullInt64
func NewNullInt64(n int64, valid bool) NullInt64 {
	return NullInt64{
		sql.NullInt64{
			Valid: valid,
			Int64: n,
		},
	}
}

// UnmarshalText initializes an invalid NullInt64 if text is empty
// otherwise it tries to parse it as an integer in base 10
func (n *NullInt64) UnmarshalText(text []byte) (err error) {
	if len(text) == 0 {
		n.Valid = false
		return nil
	}

	n.Int64, err = strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return err
	}

	n.Valid = true
	return nil
}

// UnmarshalText initializes an invalid NullInt64 if text is empty
// otherwise it tries to parse it as an integer in base 10
// MarshalText returns "" for invalid NullInt64s, otherwise the integer value
func (n NullInt64) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatInt(n.Int64, 10)), nil
}

// NullBool wraps sql.NullBool to allow it to be serializable to/from XML
// via TextMarshaler and TextUnmarshaler
type NullBool struct {
	sql.NullBool
}

// NewNullBool creats a new NullBool
func NewNullBool(b bool, valid bool) NullBool {
	return NullBool{
		sql.NullBool{
			Valid: valid,
			Bool:  b,
		},
	}
}

// UnmarshalText initializes an invalid NullBool if text is empty
// otherwise it tries to parse it as a boolean
func (n *NullBool) UnmarshalText(text []byte) (err error) {
	if len(text) == 0 {
		n.Valid = false
		return nil
	}

	n.Bool, err = strconv.ParseBool(string(text))
	if err != nil {
		return err
	}

	n.Valid = true
	return nil
}

// UnmarshalText initializes an invalid NullBool if text is empty
// otherwise it tries to parse it as an integer in base 10
// MarshalText returns "" for invalid NullBools, otherwise the boolean value
func (n NullBool) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatBool(n.Bool)), nil
}
