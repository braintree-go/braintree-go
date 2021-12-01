package braintree

import "testing"

func TestCurrencyScale(t *testing.T) {
	tests := []struct {
		name     string
		currency string
		want     uint
	}{
		{
			name:     "unknown currency",
			currency: "XXX",
			want:     2,
		},
		{
			name:     "known currency with a lower case",
			currency: "jpy",
			want:     0,
		},
		{
			name:     "currency with a scale of 2 by default",
			currency: "EUR",
			want:     2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CurrencyScale(tt.currency); got != tt.want {
				t.Errorf("CurrencyScale() = %v, want %v", got, tt.want)
			}
		})
	}
}
