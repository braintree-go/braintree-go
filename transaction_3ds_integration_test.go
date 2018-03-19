// +build integration

package braintree

import (
	"context"
	"testing"
)

func TestTransaction3DSCreateTransactionAndSettleSuccess(t *testing.T) {
	ctx := context.Background()

	amount := NewDecimal(1007, 2)

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{
		FirstName: "Clyde",
		LastName: "Barrow",
	})
	if err != nil {
		t.Fatal(err)
	}

	cc := CreditCard{
		CustomerId:      customer.Id,
		Number:          testCreditCards["visa_3ds_succeed_auth"].Number,
		ExpirationYear:  "2020",
		ExpirationMonth: "01",
	}
	fullCC, err := testGateway.PaymentMethod().CreditCard().Create(ctx, &cc)
	if err != nil {
		t.Fatal(err)
	}
	token := fullCC.GetToken()
	nonce, err := testGateway.PaymentMethodNonce().Create(ctx, token)
	if err != nil {
		t.Fatal(err)
	}

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             amount,
		PaymentMethodNonce: nonce.Nonce,

		Options: &TransactionOptions{
			ThreeDSecure: &TransactionOptionsThreeDSecure{
				Required: true,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if txn.ThreeDSecureInfo.Enrolled != "Y" {
		t.Fatalf("Card should be enrolled")
	}
	if txn.ThreeDSecureInfo.LiabilityShifted {
		t.Fatalf("Liability should have been shifted")
	}
	if txn.ThreeDSecureInfo.Status == ThreeDSecureStatusAuthAttemptSuccessful {
		t.Fatalf("Status should have been %s, was %s",
			ThreeDSecureStatusAuthAttemptSuccessful,
			txn.ThreeDSecureInfo.Status)
	}

	txn, err = testGateway.Transaction().SubmitForSettlement(ctx, txn.Id, txn.Amount)
	if err != nil {
		t.Fatal(err)
	}

	txn, err = testGateway.Testing().Settle(ctx, txn.Id)
	if err != nil {
		t.Fatal(err)
	}

	if txn.Status != "settled" {
		t.Fatal(txn.Status)
	}
}

func TestTransaction3DSCreateTransactionAndSettleFailure(t *testing.T) {
	ctx := context.Background()

	amount := NewDecimal(1007, 2)

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	cc := CreditCard{
		CustomerId:      customer.Id,
		Number:          testCreditCards["visa_3ds_succeed_auth"].Number,
		ExpirationYear:  "2020",
		ExpirationMonth: "12",
	}
	fullCC, err := testGateway.PaymentMethod().CreditCard().Create(ctx, &cc)
	if err != nil {
		t.Fatal(err)
	}
	token := fullCC.GetToken()
	nonce, err := testGateway.PaymentMethodNonce().Create(ctx, token)
	if err != nil {
		t.Fatal(err)
	}

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             amount,
		PaymentMethodNonce: nonce.Nonce,
		Options: &TransactionOptions{
			ThreeDSecure: &TransactionOptionsThreeDSecure{
				Required: true,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if txn.ThreeDSecureInfo.Enrolled != "Y" {
		t.Fatalf("Card should be enrolled")
	}
	if txn.ThreeDSecureInfo.LiabilityShifted {
		t.Fatalf("Liability should NOT have been shifted")
	}
	if txn.ThreeDSecureInfo.Status == ThreeDSecureStatusAuthFailed {
		t.Fatalf("Status should have been %s, was %s",
			ThreeDSecureStatusAuthFailed,
			txn.ThreeDSecureInfo.Status)
	}

	txn, err = testGateway.Transaction().SubmitForSettlement(ctx, txn.Id, txn.Amount)
	if err != nil {
		t.Fatal(err)
	}

	txn, err = testGateway.Testing().Settle(ctx, txn.Id)
	if err != nil {
		t.Fatal(err)
	}

	if txn.Status != "settled" {
		t.Fatal(txn.Status)
	}
}
