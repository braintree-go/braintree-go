// +build integration

package braintree

import (
	"context"
	"reflect"
	"testing"
)

func TestCustomerPayPalAccount(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	nonce := FakeNoncePayPalFuturePayment

	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: nonce,
	})
	if err != nil {
		t.Fatal(err)
	}
	paypalAccount := paymentMethod.(*PayPalAccount)

	customerFound, err := testGateway.Customer().Find(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	if customerFound.PayPalAccounts == nil || len(customerFound.PayPalAccounts.PayPalAccount) != 1 {
		t.Fatalf("Customer %#v expected to have one PayPalAccount", customerFound)
	}
	if !reflect.DeepEqual(customerFound.PayPalAccounts.PayPalAccount[0], paypalAccount) {
		t.Fatalf("Got Customer %#v PayPalAccount %#v, want %#v", customerFound, customerFound.PayPalAccounts.PayPalAccount[0], paypalAccount)
	}
}
