// +build integration

package braintree

import (
	"context"
	"testing"
)

func TestSettleTransaction(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:   "sale",
		Amount: randomAmount(),
		CreditCard: &CreditCard{
			Number:         testCardVisa,
			ExpirationDate: "05/14",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	txn, err = testGateway.Transaction().SubmitForSettlement(ctx, txn.Id, txn.Amount)
	if err != nil {
		t.Fatal(err)
	}

	prodGateway := New(Production, "my_merchant_id", "my_public_key", "my_private_key")

	_, err = prodGateway.Testing().Settle(ctx, txn.Id)
	if err.Error() != "Operation not allowed in production environment" {
		t.Log(testGateway.Environment())
		t.Fatal(err)
	}

	txn, err = testGateway.Testing().Settle(ctx, txn.Id)
	if err != nil {
		t.Fatal(err)
	}

	if txn.Status != TransactionStatusSettled {
		t.Fatal(txn.Status)
	}
}

func TestSettlementConfirmTransaction(t *testing.T) {
	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:   "sale",
		Amount: randomAmount(),
		CreditCard: &CreditCard{
			Number:         testCardVisa,
			ExpirationDate: "05/14",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	txn, err = testGateway.Transaction().SubmitForSettlement(ctx, txn.Id, txn.Amount)
	if err != nil {
		t.Fatal(err)
	}

	prodGateway := New(Production, "my_merchant_id", "my_public_key", "my_private_key")

	_, err = prodGateway.Testing().SettlementConfirm(ctx, txn.Id)
	if err.Error() != "Operation not allowed in production environment" {
		t.Log(prodGateway.Environment())
		t.Fatal(err)
	}

	txn, err = testGateway.Testing().SettlementConfirm(ctx, txn.Id)
	if err != nil {
		t.Fatal(err)
	}

	if txn.Status != TransactionStatusSettlementConfirmed {
		t.Fatal(txn.Status)
	}
}

func TestSettlementDeclinedTransaction(t *testing.T) {
	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:   "sale",
		Amount: randomAmount(),
		CreditCard: &CreditCard{
			Number:         testCardVisa,
			ExpirationDate: "05/14",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	txn, err = testGateway.Transaction().SubmitForSettlement(ctx, txn.Id, txn.Amount)
	if err != nil {
		t.Fatal(err)
	}

	prodGateway := New(Production, "my_merchant_id", "my_public_key", "my_private_key")

	_, err = prodGateway.Testing().SettlementDecline(ctx, txn.Id)
	if err.Error() != "Operation not allowed in production environment" {
		t.Log(prodGateway.Environment())
		t.Fatal(err)
	}

	txn, err = testGateway.Testing().SettlementDecline(ctx, txn.Id)
	if err != nil {
		t.Fatal(err)
	}

	if txn.Status != TransactionStatusSettlementDeclined {
		t.Fatal(txn.Status)
	}
}

func TestSettlementPendingTransaction(t *testing.T) {
	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:   "sale",
		Amount: randomAmount(),
		CreditCard: &CreditCard{
			Number:         testCardVisa,
			ExpirationDate: "05/14",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	txn, err = testGateway.Transaction().SubmitForSettlement(ctx, txn.Id, txn.Amount)
	if err != nil {
		t.Fatal(err)
	}

	prodGateway := New(Production, "my_merchant_id", "my_public_key", "my_private_key")

	_, err = prodGateway.Testing().SettlementPending(ctx, txn.Id)
	if err.Error() != "Operation not allowed in production environment" {
		t.Log(prodGateway.Environment())
		t.Fatal(err)
	}

	txn, err = testGateway.Testing().SettlementPending(ctx, txn.Id)
	if err != nil {
		t.Fatal(err)
	}

	if txn.Status != TransactionStatusSettlementPending {
		t.Fatal(txn.Status)
	}
}

func TestTransactionCreateSettleCheckCreditCardDetails(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	amount := NewDecimal(10000, 2)
	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:   "sale",
		Amount: amount,
		CreditCard: &CreditCard{
			Number:         testCardDiscover,
			ExpirationDate: "05/14",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if txn.PaymentInstrumentType != PaymentInstrumentTypeCreditCard {
		t.Fatalf("Returned payment instrument doesn't match input, expected %q, got %q",
			PaymentInstrumentTypeCreditCard, txn.PaymentInstrumentType)
	}
	if txn.CreditCard.CardType != "Discover" {
		t.Fatalf("Returned credit card detail doesn't match input, expected %q, got %q",
			"Visa", txn.CreditCard.CardType)
	}

	txn, err = testGateway.Transaction().SubmitForSettlement(ctx, txn.Id, txn.Amount)
	if err != nil {
		t.Fatal(err)
	}

	txn, err = testGateway.Testing().Settle(ctx, txn.Id)
	if err != nil {
		t.Fatal(err)
	}

	if txn.Status != TransactionStatusSettled {
		t.Fatal(txn.Status)
	}
}
