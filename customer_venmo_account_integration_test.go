// +build integration

package braintree

import (
	"context"
	"reflect"
	"testing"
)

func TestCustomerVenmoAccount(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &Customer{})
	if err != nil {
		t.Fatal(err)
	}

	nonce := FakeNonceVenmoAccount

	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: nonce,
	})
	if err != nil {
		t.Fatal(err)
	}
	venmoAccount := paymentMethod.(*VenmoAccount)

	customerFound, err := testGateway.Customer().Find(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	if customerFound.VenmoAccounts == nil || len(customerFound.VenmoAccounts.VenmoAccount) != 1 {
		t.Fatalf("Customer %#v expected to have one VenmoAccount", customerFound)
	}
	if !reflect.DeepEqual(customerFound.VenmoAccounts.VenmoAccount[0], venmoAccount) {
		t.Fatalf("Got Customer %#v VenmoAccount %#v, want %#v", customerFound, customerFound.VenmoAccounts.VenmoAccount[0], venmoAccount)
	}
}
