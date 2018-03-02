// +build integration

package braintree

import (
	"context"
	"testing"
)

func TestPayPalAccount(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cust, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	nonce := FakeNoncePayPalBillingAgreement

	g := testGateway.PayPalAccount()
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         cust.Id,
		PaymentMethodNonce: nonce,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Find
	paypalAccount, err := g.Find(ctx, paymentMethod.GetToken())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(paypalAccount)

	if paypalAccount.Token == "" {
		t.Fatal("invalid token")
	}

	// Update
	paypalAccount2, err := g.Update(ctx, &PayPalAccount{
		Token: paypalAccount.Token,
		Email: "new-email@example.com",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(paypalAccount2)

	if paypalAccount2.Token != paypalAccount.Token {
		t.Fatal("tokens do not match")
	}
	if paypalAccount2.Email != "new-email@example.com" {
		t.Fatal("paypalAccount email does not match")
	}

	// Delete
	err = g.Delete(ctx, paypalAccount2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindPayPalAccountBadData(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	paypalAccount, err := testGateway.PayPalAccount().Find(ctx, "invalid_token")

	t.Log(paypalAccount)

	if err == nil {
		t.Fatal("expected to get error because the token is invalid")
	}
}
