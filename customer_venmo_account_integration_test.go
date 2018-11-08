// +build integration

package braintree

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestCustomerVenmoAccount(t *testing.T) {
	if time.Date(2018, 11, 14, 0, 0, 0, 0, time.UTC).Sub(time.Now()) > 0 {
		t.Skip("This test started failing with a 404 when creating the payment method. It seems like an environmental issue.")
	}
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
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

	if !t.Failed() {
		t.Fatalf("This test was being temporarily skipped, but it has started passing again. The skip and this code should be removed.")
	}
}
