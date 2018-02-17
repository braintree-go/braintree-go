// +build integration

package braintree

import (
	"context"
	"reflect"
	"testing"
)

func TestCustomerAndroidPayCard(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	nonce := FakeNonceAndroidPay

	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: nonce,
	})
	if err != nil {
		t.Fatal(err)
	}
	androidPayCard := paymentMethod.(*AndroidPayCard)

	customerFound, err := testGateway.Customer().Find(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	if customerFound.AndroidPayCards == nil || len(customerFound.AndroidPayCards.AndroidPayCard) != 1 {
		t.Fatalf("Customer %#v expected to have one AndroidPayCard", customerFound)
	}
	if !reflect.DeepEqual(customerFound.AndroidPayCards.AndroidPayCard[0], androidPayCard) {
		t.Fatalf("Got Customer %#v AndroidPayCard %#v, want %#v", customerFound, customerFound.AndroidPayCards.AndroidPayCard[0], androidPayCard)
	}
}
