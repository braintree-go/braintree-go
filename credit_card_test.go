// +build unit

package braintree

import (
	"encoding/xml"
	"testing"
)

func TestCreditCardOptions_MarshalXML(t *testing.T) {
	bTrue := true
	bFalse := false
	tests := []struct {
		name    string
		cco     *CreditCardOptions
		wantXML string
		wantErr bool
	}{
		{
			name:    "nil pointer",
			cco:     nil,
			wantXML: ``,
			wantErr: false,
		},
		{
			name: "VerifyCard nil",
			cco: &CreditCardOptions{
				FailOnDuplicatePaymentMethod: true,
				MakeDefault:                  true,
			},
			wantXML: `<CreditCardOptions>
	<make-default>true</make-default>
	<fail-on-duplicate-payment-method>true</fail-on-duplicate-payment-method>
</CreditCardOptions>`,
			wantErr: false,
		},
		{
			name: "VerifyCard true",
			cco: &CreditCardOptions{
				FailOnDuplicatePaymentMethod: true,
				MakeDefault:                  true,
				VerifyCard:                   &bTrue,
			},
			wantXML: `<CreditCardOptions>
	<verify-card>true</verify-card>
	<make-default>true</make-default>
	<fail-on-duplicate-payment-method>true</fail-on-duplicate-payment-method>
</CreditCardOptions>`,
			wantErr: false,
		},
		{
			name: "VerifyCard false",
			cco: &CreditCardOptions{
				FailOnDuplicatePaymentMethod: true,
				MakeDefault:                  true,
				VerifyCard:                   &bFalse,
			},
			wantXML: `<CreditCardOptions>
	<verify-card>false</verify-card>
	<make-default>true</make-default>
	<fail-on-duplicate-payment-method>true</fail-on-duplicate-payment-method>
</CreditCardOptions>`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			output, err := xml.MarshalIndent(tt.cco, "", "\t")
			xml := string(output)
			if err != nil {
				t.Fatalf("got error = %v", err)
			}
			if xml != tt.wantXML {
				t.Errorf("got xml:\n%s\nwant xml:\n%s", xml, tt.wantXML)
			}
		})
	}
}
