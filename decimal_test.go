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
		in  *Decimal
		out []byte
	}{
		{NewDecimal(250, -2), []byte("25000")},
		{NewDecimal(2, 0), []byte("2")},
		{NewDecimal(0, 2), []byte("0.00")},
		{NewDecimal(5, 2), []byte("0.05")},
		{NewDecimal(250, 2), []byte("2.50")},
		{NewDecimal(4586, 2), []byte("45.86")},
		{NewDecimal(-5504, 2), []byte("-55.04")},
	}

	for _, tt := range tests {
		b, err := tt.in.MarshalText()
		if err != nil {
			t.Errorf("expected %+v.MarshalText() => to not error, but it did with %s", tt.in, err)
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

func TestNullDecimalUnmarshalText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in          []byte
		out         *NullDecimal
		shouldError bool
	}{
		{[]byte("2.50"), NewNullDecimal(250, 2), false},
		{[]byte("2"), NewNullDecimal(2, 0), false},
		{[]byte("0.00"), NewNullDecimal(0, 2), false},
		{[]byte("-5.504"), NewNullDecimal(-5504, 3), false},
		{[]byte("0.5"), NewNullDecimal(5, 1), false},
		{[]byte(".5"), NewNullDecimal(5, 1), false},
		{[]byte("5.504.98"), NewNullDecimal(0, 0), true},
		{[]byte("5E6"), NewNullDecimal(0, 0), true},
		{[]byte(""), &NullDecimal{}, false},
	}

	for _, tt := range tests {
		d := &NullDecimal{}
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

func TestNullDecimalMarshalText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  *NullDecimal
		out []byte
	}{
		{NewNullDecimal(250, -2), []byte("25000")},
		{NewNullDecimal(2, 0), []byte("2")},
		{NewNullDecimal(0, 2), []byte("0.00")},
		{NewNullDecimal(5, 2), []byte("0.05")},
		{NewNullDecimal(250, 2), []byte("2.50")},
		{NewNullDecimal(4586, 2), []byte("45.86")},
		{NewNullDecimal(-5504, 2), []byte("-55.04")},
		{&NullDecimal{}, []byte("")},
	}

	for _, tt := range tests {
		b, err := tt.in.MarshalText()
		if err != nil {
			t.Errorf("expected %+v.MarshalText() => to not error, but it did with %s", tt.in, err)
		}
		if string(tt.out) != string(b) {
			t.Errorf("%+v.MarshalText() => %s, want %s", tt.in, b, tt.out)
		}
	}
}

func TestNullDecimalCmp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		x, y *NullDecimal
		out  int
	}{
		{NewNullDecimal(250, -2), NewNullDecimal(250, -2), 0},
		{NewNullDecimal(2, 0), NewNullDecimal(250, -2), -1},
		{NewNullDecimal(500, 2), NewNullDecimal(50, 1), 0},
		{NewNullDecimal(2500, -2), NewNullDecimal(250, -2), 1},
		{NewNullDecimal(100, 2), NewNullDecimal(1, 0), 0},
	}

	for i, tt := range tests {
		if out := tt.x.Cmp(tt.y); out != tt.out {
			t.Errorf("%d: %+v.Cmp(%+v) => %d, want %d", i, tt.x, tt.y, out, tt.out)
		}
	}
}
