package braintree

import (
	"testing"
  "braintree/test"
)

var (
	testConfiguration = Configuration{
		environment: Sandbox,
		merchantId:  "4ngqq224rnk6gvxh",
		publicKey:   "jkq28pcxj4r85dwr",
		privateKey:  "66062a3876e2dc298f2195f0bf173f5a",
	}

	testGateway = NewGateway(testConfiguration)
)

func Test_Transaction_Create(t *testing.T) {
	tx := Transaction{
		Amount: 100,
		CreditCard: CreditCard{
			Number:         test.CreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}
	request := NewTransactionRequest(tx)

	response := testGateway.Transaction().Sale(request)

	if !response.IsSuccess() {
		t.Errorf("Transaction create response was unsuccessful")
	}
}
