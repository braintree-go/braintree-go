package braintree

import (
	"bytes"
	"testing"
)

func TestNullInt64UnmarshalText(t *testing.T) {
	tests := []struct {
		in          []byte
		out         NullInt64
		sholudError bool
	}{
		{[]byte(""), NewNullInt64(0, false), false},
		{[]byte("10"), NewNullInt64(10, true), false},
		{[]byte("abcd"), NewNullInt64(0, false), true},
	}

	for _, tt := range tests {
		n := NullInt64{}
		err := n.UnmarshalText(tt.in)

		if tt.sholudError {
			if err == nil {
				t.Errorf("expected UnmarshalText(%q) => to error, but it did not", tt.in)
			}
		} else {
			if err != nil {
				t.Errorf("expected UnmarshalText(%q) => to not error, but it did with %s", tt.in, err)
			}
		}

		if n != tt.out {
			t.Errorf("UnmarshalText(%q) => %q, want %q", tt.in, n, tt.out)
		}
	}
}

func TestNullInt64MarshalText(t *testing.T) {
	tests := []struct {
		in  NullInt64
		out []byte
	}{
		{NewNullInt64(0, false), []byte("")},
		{NewNullInt64(10, true), []byte("10")},
	}

	for _, tt := range tests {
		b, err := tt.in.MarshalText()

		if !bytes.Equal(b, tt.out) || err != nil {
			t.Errorf("%q.MarshalText() => (%s, %s), want (%s, %s)", tt.in, b, err, tt.out, nil)
		}
	}
}

func TestNullFloat64UnmarshalText(t *testing.T) {
	tests := []struct {
		in          []byte
		out         NullFloat64
		sholudError bool
	}{
		{[]byte(""), NewNullFloat64(0, false), false},
		{[]byte("10"), NewNullFloat64(10, true), false},
		{[]byte("abcd"), NewNullFloat64(0, false), true},
	}

	for _, tt := range tests {
		n := NullFloat64{}
		err := n.UnmarshalText(tt.in)

		if tt.sholudError {
			if err == nil {
				t.Errorf("expected UnmarshalText(%q) => to error, but it did not", tt.in)
			}
		} else {
			if err != nil {
				t.Errorf("expected UnmarshalText(%q) => to not error, but it did with %s", tt.in, err)
			}
		}

		if n != tt.out {
			t.Errorf("UnmarshalText(%q) => %q, want %q", tt.in, n, tt.out)
		}
	}
}

func TestNullFloat64MarshalText(t *testing.T) {
	tests := []struct {
		in  NullFloat64
		out []byte
	}{
		{NewNullFloat64(0, false), []byte("")},
		{NewNullFloat64(10, true), []byte("10")},
	}

	for _, tt := range tests {
		b, err := tt.in.MarshalText()

		if !bytes.Equal(b, tt.out) || err != nil {
			t.Errorf("%q.MarshalText() => (%s, %s), want (%s, %s)", tt.in, b, err, tt.out, nil)
		}
	}
}

func TestNullBoolUnmarshalText(t *testing.T) {
	tests := []struct {
		in          []byte
		out         NullBool
		sholudError bool
	}{
		{[]byte(""), NewNullBool(false, false), false},
		{[]byte("true"), NewNullBool(true, true), false},
		{[]byte("abcd"), NewNullBool(false, false), true},
	}

	for _, tt := range tests {
		n := NullBool{}
		err := n.UnmarshalText(tt.in)

		if tt.sholudError {
			if err == nil {
				t.Errorf("expected UnmarshalText(%q) => to error, but it did not", tt.in)
			}
		} else {
			if err != nil {
				t.Errorf("expected UnmarshalText(%q) => to not error, but it did with %s", tt.in, err)
			}
		}

		if n != tt.out {
			t.Errorf("UnmarshalText(%q) => %q, want %q", tt.in, n, tt.out)
		}
	}
}

func TestNullBoolMarshalText(t *testing.T) {
	tests := []struct {
		in  NullBool
		out []byte
	}{
		{NewNullBool(false, false), []byte("")},
		{NewNullBool(true, true), []byte("true")},
	}

	for _, tt := range tests {
		b, err := tt.in.MarshalText()

		if !bytes.Equal(b, tt.out) || err != nil {
			t.Errorf("%q.MarshalText() => (%s, %s), want (%s, %s)", tt.in, b, err, tt.out, nil)
		}
	}
}
