package braintree

import (
	"testing"
)

func TestTransactionCreate(t *testing.T) {
	tx := Transaction{
		Type:   "sale",
		Amount: 100.00,
		CreditCard: &CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
		Options: &TransactionOptions{
			SubmitForSettlement: true,
		},
	}

	result, err := txGateway.Sale(tx)

	if err != nil {
		t.Errorf(err.Error())
	} else if !result.Success() {
		t.Errorf("Transaction create response was unsuccessful")
	} else if result.Transaction().Id == "" {
		t.Errorf("Transaction did not receive an ID")
	} else if result.Transaction().Status != "submitted_for_settlement" {
		t.Errorf("Transaction was not submitted for settlement")
	}
}

func TestTransactionCreateWhenGatewayRejected(t *testing.T) {
	tx := Transaction{
		Type:   "sale",
		Amount: 2010.00,
		CreditCard: &CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}

	result, err := txGateway.Sale(tx)

	if err != nil {
		t.Errorf(err.Error())
	} else if result.Success() {
		t.Errorf("Transaction create response was successful, expected failure")
	} else if result.Message() != "Card Issuer Declined CVV" {
		t.Errorf("Got wrong error message. Got: " + result.Message())
	}
}

func TestFindTransaction(t *testing.T) {
	tx := Transaction{
		Type:   "sale",
		Amount: 100.00,
		CreditCard: &CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}

	saleResult, err := txGateway.Sale(tx)

	if err != nil {
		t.Errorf(err.Error())
	} else if !saleResult.Success() {
		t.Errorf("Transaction create response was unsuccessful")
	}

	txId := saleResult.Transaction().Id

	findResult, err := txGateway.Find(txId)

	if err != nil {
		t.Errorf(err.Error())
	} else if !findResult.Success() {
		t.Errorf("Transaction find response was unsuccessful")
	} else if findResult.Transaction().Id != txId {
		t.Errorf("Transaction find came back with the wrong transaction!")
	}
}

func TestFindNonExistantTransaction(t *testing.T) {
	result, err := txGateway.Find("bad transaction ID")

	if err == nil {
		t.Errorf("Did not receive an error when trying to find a non-existant transaction")
	} else if result.Success() {
		t.Errorf("Transaction find response was successful on bad data")
	} else if err.Error() != "A transaction with that ID could not be found" {
		t.Errorf("Got the wrong error message when finding a non-existant transaction. Got: " + err.Error())
	}
}
