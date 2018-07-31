// +build integration

package braintree

import (
	"context"
	"testing"
)

func TestTransaction3DSRequiredGatewayRejected(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	amount := NewDecimal(1007, 2)

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	cc, err := testGateway.PaymentMethod().CreditCard().Create(ctx, &CreditCard{
		CustomerId:      customer.Id,
		Number:          testCardVisa,
		ExpirationYear:  "2020",
		ExpirationMonth: "01",
	})
	if err != nil {
		t.Fatal(err)
	}

	nonce, err := testGateway.PaymentMethodNonce().Create(ctx, cc.Token)
	if err != nil {
		t.Fatal(err)
	}
	if nonce.ThreeDSecureInfo != nil {
		t.Fatalf("Nonce 3DS Info present when card was non-3DS")
	}

	_, err = testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             amount,
		PaymentMethodNonce: nonce.Nonce,
		Options: &TransactionOptions{
			ThreeDSecure: &TransactionOptionsThreeDSecureRequest{Required: true},
		},
	})
	if err == nil {
		t.Fatal("Did not receive error when creating transaction requiring 3DS with non-3DS nonce")
	}
	if err.Error() != "Gateway Rejected: three_d_secure" {
		t.Fatal(err)
	}
	if err.(*BraintreeError).Transaction.ThreeDSecureInfo != nil {
		t.Fatalf("Transaction 3DS Info present when nonce for transaction was non-3DS")
	}
}
