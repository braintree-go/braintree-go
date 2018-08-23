// +build integration

package braintree

import (
	"context"
	"testing"
)

func TestPaymentMethodNonce(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway(t).Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	paymentMethod, err := testGateway(t).PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactableVisa,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", paymentMethod)

	paymentMethodNonce, err := testGateway(t).PaymentMethodNonce().Create(ctx, paymentMethod.GetToken())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", paymentMethodNonce)
	t.Logf("%#v", paymentMethodNonce.Details)

	if paymentMethodNonce.Type != "CreditCard" {
		t.Errorf("nonce type got %q, want %q", paymentMethodNonce.Type, "CreditCard")
	}

	paymentMethodNonceFound, err := testGateway(t).PaymentMethodNonce().Find(ctx, paymentMethodNonce.Nonce)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", paymentMethodNonceFound)
	t.Logf("%#v", paymentMethodNonceFound.Details)

	if paymentMethodNonceFound.Type != "CreditCard" {
		t.Errorf("found nonce type got %q, want %q", paymentMethodNonceFound.Type, "CreditCard")
	}
	if paymentMethodNonceFound.Nonce != paymentMethodNonce.Nonce {
		t.Errorf("found nonce got %q, want %q", paymentMethodNonceFound.Nonce, paymentMethodNonce.Nonce)
	}
}
