package braintree

import (
	"testing"
)

var (
	testConfiguration = Configuration{
		environment: Sandbox,
		merchantId:  "4ngqq224rnk6gvxh",
		publicKey:   "jkq28pcxj4r85dwr",
		privateKey:  "66062a3876e2dc298f2195f0bf173f5a",
	}

	gateway = NewGateway(testConfiguration)
)

func TestTransactionCreate(t *testing.T) {
	tx := Transaction{
    Type: "sale",
		Amount: 100,
		CreditCard: CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}
	request := NewTransactionRequest(tx)

	response, _ := gateway.ExecuteTransactionRequest(request)

	if !response.IsSuccess() {
		t.Errorf("Transaction create response was unsuccessful")
	}
}
