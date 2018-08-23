// +build integration

package braintree

import (
	"context"
	"testing"
)

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestClientToken(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	g := testGateway(t).ClientToken()
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

	customer, err := testGateway(t).Customer().Create(ctx, customerRequest)
	if err != nil {
		t.Error(err)
	}

	customerId := customer.Id

	token, err := testGateway(t).ClientToken().GenerateWithCustomer(ctx, customerId)
	if err != nil {
		t.Error(err)
	} else if len(token) == 0 {
		t.Fatalf("Received empty client token")
	}
}
