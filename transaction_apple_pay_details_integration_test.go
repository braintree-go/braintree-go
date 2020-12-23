// +build integration

package braintree

import (
	"context"
	"testing"

	"github.com/braintree-go/braintree-go/testhelpers"
)

func TestTransactionApplePayDetails(t *testing.T) {
	ctx := context.Background()

	tx, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(2000, 2),
		PaymentMethodNonce: FakeNonceApplePayVisa,
	})

	t.Log(tx)

	if err != nil {
		t.Fatal(err)
	}
	if tx.Id == "" {
		t.Fatal("Received invalid ID on new transaction")
	}
	if tx.Status != TransactionStatusAuthorized {
		t.Fatal(tx.Status)
	}

	if tx.ApplePayDetails == nil {
		t.Fatal("Expected ApplePayDetails for transaction created with ApplePay nonce")
	}

	t.Log(tx.ApplePayDetails)

	wantNonceCardType := "Apple Pay - Visa"
	if tx.ApplePayDetails.CardType != wantNonceCardType {
		t.Errorf("Got ApplePayDetails.CardType %v, want %v", tx.ApplePayDetails.CardType, wantNonceCardType)
	}
	if tx.ApplePayDetails.PaymentInstrumentName == "" {
		t.Fatal("Expected ApplePayDetails to have PaymentInstrumentName set")
	}
	if tx.ApplePayDetails.SourceDescription == "" {
		t.Fatal("Expected ApplePayDetails to have SourceDescription set")
	}
	if tx.ApplePayDetails.CardholderName == "" {
		t.Fatal("Expected ApplePayDetails to have CardholderName set")
	}
	if !testhelpers.ValidExpiryMonth(tx.ApplePayDetails.ExpirationMonth) {
		t.Errorf("ApplePayDetails.ExpirationMonth (%s) does not match expected value", tx.ApplePayDetails.ExpirationMonth)
	}
	if !testhelpers.ValidExpiryYear(tx.ApplePayDetails.ExpirationYear) {
		t.Errorf("ApplePayDetails.ExpirationYear (%s) does not match expected value", tx.ApplePayDetails.ExpirationYear)
	}
	if !testhelpers.ValidBIN(tx.ApplePayDetails.BIN) {
		t.Errorf("ApplePayDetails.BIN (%s) does not conform expected value", tx.ApplePayDetails.BIN)
	}
	if !testhelpers.ValidLast4(tx.ApplePayDetails.Last4) {
		t.Errorf("ApplePayDetails.Last4 (%s) does not conform match value", tx.ApplePayDetails.Last4)
	}
}

func TestTransactionWithoutApplePayDetails(t *testing.T) {
	ctx := context.Background()

	tx, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(2000, 2),
		PaymentMethodNonce: FakeNonceTransactable,
	})

	t.Log(tx)

	if err != nil {
		t.Fatal(err)
	}
	if tx.Id == "" {
		t.Fatal("Received invalid ID on new transaction")
	}
	if tx.Status != TransactionStatusAuthorized {
		t.Fatal(tx.Status)
	}

	if tx.ApplePayDetails != nil {
		t.Fatalf("Expected ApplePayDetails to be nil for transaction created without ApplePay, but was %#v", tx.ApplePayDetails)
	}
}
