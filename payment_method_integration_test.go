// +build integration

package braintree

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/braintree-go/braintree-go/testhelpers"
)

func TestPaymentMethod(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cust, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.PaymentMethod()

	// Create using credit card
	paymentMethod, err := g.Create(ctx, &PaymentMethodRequest{
		CustomerId:         cust.Id,
		PaymentMethodNonce: FakeNonceTransactableVisa,
	})
	if err != nil {
		t.Fatal(err)
	}

	if paymentMethod.GetCustomerId() != cust.Id {
		t.Errorf("Got paymentMethod customer Id %#v, want %#v", paymentMethod.GetCustomerId(), cust.Id)
	}
	if paymentMethod.GetToken() == "" {
		t.Errorf("Got paymentMethod token %#v, want a value", paymentMethod.GetToken())
	}

	// Update using different credit card
	rand.Seed(time.Now().UTC().UnixNano())
	token := fmt.Sprintf("btgo_test_token_%d", rand.Int()+1)
	paymentMethod, err = g.Update(ctx, paymentMethod.GetToken(), &PaymentMethodRequest{
		PaymentMethodNonce: FakeNonceTransactableMasterCard,
		Token:              token,
	})
	if err != nil {
		t.Fatal(err)
	}

	if paymentMethod.GetToken() != token {
		t.Errorf("Got paymentMethod token %#v, want %#v", paymentMethod.GetToken(), token)
	}

	// Updating with different payment method type should fail
	if _, err = g.Update(ctx, token, &PaymentMethodRequest{PaymentMethodNonce: FakeNoncePayPalBillingAgreement}); err == nil {
		t.Errorf("Updating with a different payment method type should have failed")
	}

	// Find credit card
	paymentMethod, err = g.Find(ctx, token)
	if err != nil {
		t.Fatal(err)
	}

	if paymentMethod.GetCustomerId() != cust.Id {
		t.Errorf("Got paymentMethod customer Id %#v, want %#v", paymentMethod.GetCustomerId(), cust.Id)
	}
	if paymentMethod.GetToken() != token {
		t.Errorf("Got paymentMethod token %#v, want %#v", paymentMethod.GetToken(), token)
	}

	// Delete credit card
	if err := g.Delete(ctx, token); err != nil {
		t.Fatal(err)
	}

	// Create using PayPal
	paymentMethod, err = g.Create(ctx, &PaymentMethodRequest{
		CustomerId:         cust.Id,
		PaymentMethodNonce: FakeNoncePayPalBillingAgreement,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Find PayPal
	_, err = g.Find(ctx, paymentMethod.GetToken())
	if err != nil {
		t.Fatal(err)
	}

	// Updating a PayPal account with a different payment method nonce of any kind should fail
	if _, err = g.Update(ctx, paymentMethod.GetToken(), &PaymentMethodRequest{PaymentMethodNonce: FakeNoncePayPalOneTimePayment}); err == nil {
		t.Errorf("Updating a PayPal account with a different nonce should have failed")
	}

	// Delete PayPal
	if err := g.Delete(ctx, paymentMethod.GetToken()); err != nil {
		t.Fatal(err)
	}

	// Cleanup
	if err := testGateway.Customer().Delete(ctx, cust.Id); err != nil {
		t.Fatal(err)
	}
}

func TestPaymentMethodFailedAutoVerification(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cust, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.PaymentMethod()
	pm, err := g.Create(ctx, &PaymentMethodRequest{
		CustomerId:         cust.Id,
		PaymentMethodNonce: FakeNonceProcessorDeclinedVisa,
	})
	if err == nil {
		t.Fatal("Got no error, want error")
	}
	if g, w := err.(*BraintreeError).ErrorMessage, "Do Not Honor"; g != w {
		t.Fatalf("Got error %q, want error %q", g, w)
	}

	t.Logf("%#v\n", err)
	t.Logf("%#v\n", pm)
}

func TestPaymentMethodForceNotVerified(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cust, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.PaymentMethod()
	pm, err := g.Create(ctx, &PaymentMethodRequest{
		CustomerId:         cust.Id,
		PaymentMethodNonce: FakeNonceProcessorDeclinedVisa,
		Options: &PaymentMethodRequestOptions{
			VerifyCard: testhelpers.BoolPtr(false),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", pm)
}
