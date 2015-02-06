package braintree

import (
	"bytes"
	"strconv"
	"strings"
)

const precision = 16

// Decimal represents fixed precision numbers
type Decimal struct {
	Unscaled int64
	Scale    int
}

// NewDecimal creates a new decimal number equal to
// unscaled ** 10 ^ (-scale)
func NewDecimal(unscaled int64, scale int) *Decimal {
	return &Decimal{Unscaled: unscaled, Scale: scale}
}

// MarshalText outputs a decimal representation of the scaled number
func (d *Decimal) MarshalText() (text []byte, err error) {
	b := new(bytes.Buffer)
	if d.Scale <= 0 {
		b.WriteString(strconv.FormatInt(d.Unscaled, 10))
		b.WriteString(strings.Repeat("0", -d.Scale))
	} else {
		str := strconv.FormatInt(d.Unscaled, 10)
		b.WriteString(str[:len(str)-d.Scale])
		b.WriteString(".")
		b.WriteString(str[len(str)-d.Scale:])
	}
	return b.Bytes(), nil
}

// UnmarshalText creates a Decimal from a string representation (e.g. 5.20)
// Currently only supports decimal strings
func (d *Decimal) UnmarshalText(text []byte) (err error) {
	var (
		str            = string(text)
		unscaled int64 = 0
		scale    int   = 0
	)

	if i := strings.Index(str, "."); i != -1 {
		scale = len(str) - i - 1
		str = strings.Replace(str, ".", "", 1)
	}

	if unscaled, err = strconv.ParseInt(str, 10, 64); err != nil {
		return err
	}

	d.Unscaled = unscaled
	d.Scale = scale

	return nil
}
