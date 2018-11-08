// +build integration

package braintree

import (
	"context"
	"testing"

	"github.com/braintree-go/braintree-go/testhelpers"
)

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestClientToken(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	g := testGateway.ClientToken()
	token, err := g.Generate(ctx)
	if err != nil {
		t.Fatalf("failed to generate client token: %s", err)
	}
	if len(token) == 0 {
		t.Fatalf("empty client token!")
	}
}

func TestClientTokenWithCustomer(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customerRequest := &CustomerRequest{FirstName: "Lionel"}

	customer, err := testGateway.Customer().Create(ctx, customerRequest)
	if err != nil {
		t.Error(err)
	}

	customerId := customer.Id

	token, err := testGateway.ClientToken().GenerateWithCustomer(ctx, customerId)
	if err != nil {
		t.Error(err)
	} else if len(token) == 0 {
		t.Fatalf("Received empty client token")
	}
}

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestClientTokenGateway_GenerateWithRequest(t *testing.T) {
	// Getting customer from the API
	customer, err := testGateway.Customer().Create(context.Background(), &CustomerRequest{FirstName: "Lionel"})
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name    string
		req     *ClientTokenRequest
		wantErr bool
	}{
		{
			name: "empty request struct",
			req:  nil,
		},
		{
			name: "request with provided version",
			req:  &ClientTokenRequest{Version: 2},
		},
		{
			name:    "request with non existent customerID",
			req:     &ClientTokenRequest{CustomerID: "///////@@@@@@@"},
			wantErr: true,
		},
		{
			name: "request with customer id",
			req:  &ClientTokenRequest{CustomerID: customer.Id},
		},
		{
			name: "request with merchant id",
			req:  &ClientTokenRequest{MerchantAccountID: testMerchantAccountId},
		},
		{
			name: "request with customer id and merchant id",
			req:  &ClientTokenRequest{CustomerID: customer.Id, MerchantAccountID: testMerchantAccountId},
		},
		{
			name: "request with customer id and merchant id and options verify card not set",
			req: &ClientTokenRequest{
				CustomerID:        customer.Id,
				MerchantAccountID: testMerchantAccountId,
				Options: &ClientTokenRequestOptions{
					FailOnDuplicatePaymentMethod: true,
					MakeDefault:                  true,
				},
			},
		},
		{
			name: "request with customer id and merchant id and options verify card true",
			req: &ClientTokenRequest{
				CustomerID:        customer.Id,
				MerchantAccountID: testMerchantAccountId,
				Options: &ClientTokenRequestOptions{
					FailOnDuplicatePaymentMethod: true,
					MakeDefault:                  true,
					VerifyCard:                   testhelpers.BoolPtr(true),
				},
			},
		},
		{
			name: "request with customer id and merchant id and options verify card false",
			req: &ClientTokenRequest{
				CustomerID:        customer.Id,
				MerchantAccountID: testMerchantAccountId,
				Options: &ClientTokenRequestOptions{
					FailOnDuplicatePaymentMethod: true,
					MakeDefault:                  true,
					VerifyCard:                   testhelpers.BoolPtr(false),
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := testGateway.ClientToken()
			token, err := g.GenerateWithRequest(context.TODO(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientTokenGateway.Generate() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(token) == 0 {
				t.Errorf("ClientTokenGateway.Generate() returned empty client token!")
			}
		})
	}
}
