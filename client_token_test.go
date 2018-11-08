// +build unit

package braintree

import (
	"encoding/xml"
	"testing"

	"github.com/braintree-go/braintree-go/testhelpers"
)

func TestClientToken_MarshalXML(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		req     *ClientTokenRequest
		wantXML string
	}{
		{
			name:    "request nil",
			req:     nil,
			wantXML: ``,
		},
		{
			name: "request empty",
			req:  &ClientTokenRequest{},
			wantXML: `<client-token>
  <version>0</version>
</client-token>`,
		},
		{
			name: "request with provided version",
			req:  &ClientTokenRequest{Version: 2},
			wantXML: `<client-token>
  <version>2</version>
</client-token>`,
		},
		{
			name: "request with customer and merchant account",
			req: &ClientTokenRequest{
				CustomerID:        "1234",
				MerchantAccountID: "5678",
			},
			wantXML: `<client-token>
  <customer-id>1234</customer-id>
  <merchant-account-id>5678</merchant-account-id>
  <version>0</version>
</client-token>`,
		},
		{
			name: "request with non-pointer options false",
			req: &ClientTokenRequest{
				Options: &ClientTokenRequestOptions{
					FailOnDuplicatePaymentMethod: false,
					MakeDefault:                  false,
				},
			},
			wantXML: `<client-token>
  <options></options>
  <version>0</version>
</client-token>`,
		},
		{
			name: "request with non-pointer options true",
			req: &ClientTokenRequest{
				Options: &ClientTokenRequestOptions{
					FailOnDuplicatePaymentMethod: true,
					MakeDefault:                  true,
				},
			},
			wantXML: `<client-token>
  <options>
    <fail-on-duplicate-payment-method>true</fail-on-duplicate-payment-method>
    <make-default>true</make-default>
  </options>
  <version>0</version>
</client-token>`,
		},
		{
			name: "request with verify card true",
			req: &ClientTokenRequest{
				Options: &ClientTokenRequestOptions{
					VerifyCard: testhelpers.BoolPtr(true),
				},
			},
			wantXML: `<client-token>
  <options>
    <verify-card>true</verify-card>
  </options>
  <version>0</version>
</client-token>`,
		},
		{
			name: "request with verify card false",
			req: &ClientTokenRequest{
				Options: &ClientTokenRequestOptions{
					VerifyCard: testhelpers.BoolPtr(false),
				},
			},
			wantXML: `<client-token>
  <options>
    <verify-card>false</verify-card>
  </options>
  <version>0</version>
</client-token>`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := xml.MarshalIndent(test.req, "", "  ")
			xml := string(output)
			if err != nil {
				t.Fatalf("got error = %v", err)
			}
			if xml != test.wantXML {
				t.Errorf("got xml:\n%s\nwant xml:\n%s", xml, test.wantXML)
			}
		})
	}
}
