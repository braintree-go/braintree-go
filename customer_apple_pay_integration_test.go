// +build integration

package braintree

import (
	"context"
	"reflect"
	"testing"
)

func TestCustomerApplePayCard(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway(t).Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	nonce := FakeNonceApplePayVisa

	paymentMethod, err := testGateway(t).PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: nonce,
	})
	if err != nil {
		t.Fatal(err)
	}
	applePayCard := paymentMethod.(*ApplePayCard)

	customerFound, err := testGateway(t).Customer().Find(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	if customerFound.ApplePayCards == nil || len(customerFound.ApplePayCards.ApplePayCard) != 1 {
		t.Fatalf("Customer %#v expected to have one ApplePayCard", customerFound)
	}
	if !reflect.DeepEqual(customerFound.ApplePayCards.ApplePayCard[0], applePayCard) {
		t.Fatalf("Got Customer %#v ApplePayCard %#v, want %#v", customerFound, customerFound.ApplePayCards.ApplePayCard[0], applePayCard)
	}
}
