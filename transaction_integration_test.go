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

	gateway = NewGateway(testConfiguration)
)

func TestTransactionCreate(t *testing.T) {
	tx := Transaction{
		Type:   "sale",
		Amount: 100,
		CreditCard: CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}
	request := NewTransactionRequest(tx)

	response, err := gateway.ExecuteTransactionRequest(request)

	if err != nil {
		t.Errorf(err.Error())
	} else if !response.Success {
		t.Errorf("Transaction create response was unsuccessful")
	}
}
