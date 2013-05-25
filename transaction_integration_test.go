package braintree

import (
	"testing"
)

var (
	testConfiguration = Configuration{
		environment: Development,
		merchantId:  "integration_merchant_id",
		publicKey:   "b6fkbfmhnjdg7333",
		privateKey:  "37912780851d0f68c267ea049cfa0114",
	}

	gateway = TransactionGateway{NewGateway(testConfiguration)}
)

func TestTransactionCreate(t *testing.T) {
	tx := Transaction{
		Type:   "sale",
		Amount: 100.00,
		CreditCard: CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}

	response, err := gateway.Sale(tx)

	if err != nil {
		t.Errorf(err.Error())
	} else if !response.IsSuccess() {
		t.Errorf("Transaction create response was unsuccessful")
	} else if response.Transaction().Id == "" {
		t.Errorf("Transaction did not receive an ID")
	}
}

func TestTransactionCreateWhenGatewayRejected(t *testing.T) {
	tx := Transaction{
		Type:   "sale",
		Amount: 2010.00,
		CreditCard: CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}

	response, err := gateway.Sale(tx)

	if err != nil {
		t.Errorf(err.Error())
	} else if response.IsSuccess() {
		t.Errorf("Transaction create response was successful, expected failure")
	} else if response.GetMessage() != "Card Issuer Declined CVV" {
		t.Errorf("Got wrong error message. Got: " + response.GetMessage())
	}
}
