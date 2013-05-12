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

	testGateway = NewGateway(testConfiguration)
)

func Test_Transaction_Create(t *testing.T) {
	request := NewTransactionRequest().Amount(100).CreditCard().Number(TestCards["visa"].Number).ExpirationDate("05/2014").Done()
	response := testGateway.Transaction().Sale(request)
	if !response.IsSuccess() {
		t.Errorf("Transaction create response was unsuccessful")
	}
}
