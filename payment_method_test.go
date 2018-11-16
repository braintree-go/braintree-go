// +build unit

package braintree

import (
	"encoding/xml"
	"testing"
)

func TestPaymentMethodRequestOptions_MarshalXML(t *testing.T) {
	bTrue := true
	bFalse := false
	tests := []struct {
		name    string
		pmo     *PaymentMethodRequestOptions
		wantXML string
		wantErr bool
	}{
		{
			name:    "nil pointer",
			pmo:     nil,
			wantXML: ``,
			wantErr: false,
		},
		{
			name: "VerifyCard nil",
			pmo: &PaymentMethodRequestOptions{
				FailOnDuplicatePaymentMethod: true,
				MakeDefault:                  true,
			},
			wantXML: `<PaymentMethodRequestOptions>
	<make-default>true</make-default>
	<fail-on-duplicate-payment-method>true</fail-on-duplicate-payment-method>
</PaymentMethodRequestOptions>`,
			wantErr: false,
		},
		{
			name: "VerifyCard true",
			pmo: &PaymentMethodRequestOptions{
				FailOnDuplicatePaymentMethod: true,
				MakeDefault:                  true,
				VerifyCard:                   &bTrue,
			},
			wantXML: `<PaymentMethodRequestOptions>
	<make-default>true</make-default>
	<fail-on-duplicate-payment-method>true</fail-on-duplicate-payment-method>
	<verify-card>true</verify-card>
</PaymentMethodRequestOptions>`,
			wantErr: false,
		},
		{
			name: "VerifyCard false",
			pmo: &PaymentMethodRequestOptions{
				FailOnDuplicatePaymentMethod: true,
				MakeDefault:                  true,
				VerifyCard:                   &bFalse,
			},
			wantXML: `<PaymentMethodRequestOptions>
	<make-default>true</make-default>
	<fail-on-duplicate-payment-method>true</fail-on-duplicate-payment-method>
	<verify-card>false</verify-card>
</PaymentMethodRequestOptions>`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			output, err := xml.MarshalIndent(tt.pmo, "", "\t")
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

func TestClientTokenRequestOptions_MarshalXML(t *testing.T) {
	bTrue := true
	bFalse := false
	tests := []struct {
		name    string
		ctro    *ClientTokenRequestOptions
		wantXML string
		wantErr bool
	}{
		{
			name:    "nil pointer",
			ctro:    nil,
			wantXML: ``,
			wantErr: false,
		},
		{
			name: "VerifyCard nil",
			ctro: &ClientTokenRequestOptions{
				FailOnDuplicatePaymentMethod: true,
				MakeDefault:                  true,
			},
			wantXML: `<ClientTokenRequestOptions>
	<fail-on-duplicate-payment-method>true</fail-on-duplicate-payment-method>
	<make-default>true</make-default>
</ClientTokenRequestOptions>`,
			wantErr: false,
		},
		{
			name: "VerifyCard true",
			ctro: &ClientTokenRequestOptions{
				FailOnDuplicatePaymentMethod: true,
				MakeDefault:                  true,
				VerifyCard:                   &bTrue,
			},
			wantXML: `<ClientTokenRequestOptions>
	<fail-on-duplicate-payment-method>true</fail-on-duplicate-payment-method>
	<make-default>true</make-default>
	<verify-card>true</verify-card>
</ClientTokenRequestOptions>`,
			wantErr: false,
		},
		{
			name: "VerifyCard false",
			ctro: &ClientTokenRequestOptions{
				FailOnDuplicatePaymentMethod: true,
				MakeDefault:                  true,
				VerifyCard:                   &bFalse,
			},
			wantXML: `<ClientTokenRequestOptions>
	<fail-on-duplicate-payment-method>true</fail-on-duplicate-payment-method>
	<make-default>true</make-default>
	<verify-card>false</verify-card>
</ClientTokenRequestOptions>`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			output, err := xml.MarshalIndent(tt.ctro, "", "\t")
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
