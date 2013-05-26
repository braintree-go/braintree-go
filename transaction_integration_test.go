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

	response, err := txGateway.Sale(tx)

	if err != nil {
		t.Errorf(err.Error())
	} else if !response.IsSuccess() {
		t.Errorf("Transaction create response was unsuccessful")
	} else if response.Transaction().Id == "" {
		t.Errorf("Transaction did not receive an ID")
	} else if response.Transaction().Status != "submitted_for_settlement" {
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

	response, err := txGateway.Sale(tx)

	if err != nil {
		t.Errorf(err.Error())
	} else if response.IsSuccess() {
		t.Errorf("Transaction create response was successful, expected failure")
	} else if response.GetMessage() != "Card Issuer Declined CVV" {
		t.Errorf("Got wrong error message. Got: " + response.GetMessage())
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

	response, err := txGateway.Sale(tx)

	if err != nil {
		t.Errorf(err.Error())
	} else if !response.IsSuccess() {
		t.Errorf("Transaction create response was unsuccessful")
	}

	txId := response.Transaction().Id

	response, err = txGateway.Find(txId)

	if err != nil {
		t.Errorf(err.Error())
	} else if !response.IsSuccess() {
		t.Errorf("Transaction find response was unsuccessful")
	} else if response.Transaction().Id != txId {
		t.Errorf("Transaction find came back with the wrong transaction!")
	}
}

func TestFindNonExistantTransaction(t *testing.T) {
	response, err := txGateway.Find("bad transaction ID")

	if err == nil {
		t.Errorf("Did not receive an error when trying to find a non-existant transaction")
	} else if response.IsSuccess() {
		t.Errorf("Transaction find response was successful on bad data")
	} else if err.Error() != "A transaction with that ID could not be found" {
		t.Errorf("Got the wrong error message when finding a non-existant transaction")
	}
}
