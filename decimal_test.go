//go:build unit
// +build unit

package braintree

import (
	"reflect"
	"testing"
)

func TestDecimalUnmarshalText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in          []byte
		out         *Decimal
		shouldError bool
	}{
		{[]byte("2.50"), NewDecimal(250, 2), false},
		{[]byte("2"), NewDecimal(2, 0), false},
		{[]byte("0.00"), NewDecimal(0, 2), false},
		{[]byte("-5.504"), NewDecimal(-5504, 3), false},
		{[]byte("0.5"), NewDecimal(5, 1), false},
		{[]byte(".5"), NewDecimal(5, 1), false},
		{[]byte("5.504.98"), NewDecimal(0, 0), true},
		{[]byte("5E6"), NewDecimal(0, 0), true},
	}

	for _, tt := range tests {
		d := &Decimal{}
		err := d.UnmarshalText(tt.in)

		if tt.shouldError {
			if err == nil {
				t.Errorf("expected UnmarshalText(%s) => to error, but it did not", tt.in)
			}
		} else {
			if err != nil {
				t.Errorf("expected UnmarshalText(%s) => to not error, but it did with %s", tt.in, err)
			}
		}

		if !reflect.DeepEqual(d, tt.out) {
			t.Errorf("UnmarshalText(%s) => %+v, want %+v", tt.in, d, tt.out)
		}
	}
}

func TestDecimalMarshalText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in          *Decimal
		out         []byte
		shouldError bool
	}{
		{NewDecimal(250, -2), []byte("25000"), false},
		{NewDecimal(2, 0), []byte("2"), false},
		{NewDecimal(23, 0), []byte("23"), false},
		{NewDecimal(234, 0), []byte("234"), false},
		{NewDecimal(0, 1), []byte("0.0"), false},
		{NewDecimal(1, 1), []byte("0.1"), false},
		{NewDecimal(12, 1), []byte("1.2"), false},
		{NewDecimal(0, 2), []byte("0.00"), false},
		{NewDecimal(5, 2), []byte("0.05"), false},
		{NewDecimal(55, 2), []byte("0.55"), false},
		{NewDecimal(250, 2), []byte("2.50"), false},
		{NewDecimal(4586, 2), []byte("45.86"), false},
		{NewDecimal(-5504, 2), []byte("-55.04"), false},
		{NewDecimal(0, 3), []byte("0.000"), false},
		{NewDecimal(5, 3), []byte("0.005"), false},
		{NewDecimal(55, 3), []byte("0.055"), false},
		{NewDecimal(250, 3), []byte("0.250"), false},
		{NewDecimal(4586, 3), []byte("4.586"), false},
		{NewDecimal(45867, 3), []byte("45.867"), false},
		{NewDecimal(-55043, 3), []byte("-55.043"), false},
		{nil, nil, true},
	}

	for _, tt := range tests {
		b, err := tt.in.MarshalText()

		if tt.shouldError {
			if err == nil {
				t.Errorf("expected %+v.MarshalText() => to error, but it did not", tt.in)
			}
		} else {
			if err != nil {
				t.Errorf("expected %+v.MarshalText() => to not error, but it did with %s", tt.in, err)
			}
		}

		if string(tt.out) != string(b) {
			t.Errorf("%+v.MarshalText() => %s, want %s", tt.in, b, tt.out)
		}
	}
}

func TestDecimalCmp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		x, y *Decimal
		out  int
	}{
		{NewDecimal(250, -2), NewDecimal(250, -2), 0},
		{NewDecimal(2, 0), NewDecimal(250, -2), -1},
		{NewDecimal(500, 2), NewDecimal(50, 1), 0},
		{NewDecimal(2500, -2), NewDecimal(250, -2), 1},
		{NewDecimal(100, 2), NewDecimal(1, 0), 0},
	}

	for i, tt := range tests {
		if out := tt.x.Cmp(tt.y); out != tt.out {
			t.Errorf("%d: %+v.Cmp(%+v) => %d, want %d", i, tt.x, tt.y, out, tt.out)
		}
	}
}

func TestNewDecimalByCurrency(t *testing.T) {
	tests := []struct {
		name     string
		currency string
		amount   float64
		expected *Decimal
	}{
		{
			name:     "Test EUR currency",
			currency: "EUR",
			amount:   10.50,
			expected: NewDecimal(1050, 2),
		},
		{
			name:     "Test EUR currency, zero case",
			currency: "EUR",
			amount:   0,
			expected: NewDecimal(0, 2),
		},
		{
			name:     "Test unknown (UKN) currency, default should be used",
			currency: "UKN",
			amount:   10.60,
			expected: NewDecimal(1060, 2),
		},
		{
			name:     "Test CVE currency with zero decimal adjustment",
			currency: "CVE",
			amount:   150,
			expected: NewDecimal(150, 0),
		},
		{
			name:     "Test BHD currency with 3 decimal adjustment points",
			currency: "BHD",
			amount:   150.050,
			expected: NewDecimal(150050, 3),
		},
		{
			name:     "Test JPY currency with 0 decimal adjustment points",
			currency: "JPY",
			amount:   150.020,
			expected: NewDecimal(150, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDecimalByCurrency(tt.currency, tt.amount)

			if !reflect.DeepEqual(d, tt.expected) {
				t.Errorf("%s: got %d, want %d", tt.name, d, tt.expected)
			}
		})
	}
}
