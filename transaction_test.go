package braintree

import "testing"

func TestCreateErrorHandling(t *testing.T) {
	gateway := TransactionGateway{blowUpGateway{}}
	response, err := gateway.Create(Transaction{})

	if response.Success() {
		t.Errorf("Sale response was successful when the server 500'd")
	} else if err.Error() != "Unexpected response from server: 500 Internal Server Error" {
		t.Errorf("Sale returned wrong error when the server 500'd. Got: " + err.Error())
	}
}

func TestCreateBadInputHandling(t *testing.T) {
	gateway := TransactionGateway{badInputGateway{}}
	response, err := gateway.Create(Transaction{})

	if response.Success() {
		t.Errorf("Sale response was successful when the server 422'd")
	} else if err != nil {
		t.Errorf("Sale returned an error that should be handled on response object")
	} else if response.Message() != "Card Issuer Declined CVV" {
		t.Errorf("Sale returned wrong error message")
	}
}

func TestFindErrorHandling(t *testing.T) {
	gateway := TransactionGateway{notFoundGateway{}}
	response, err := gateway.Find("bogusId")

	if response.Success() {
		t.Errorf("Find response was successful when the server 404'd")
	} else if err.Error() != "A transaction with that ID could not be found" {
		t.Errorf("Find returned wrong error when the server 404'd")
	}
}
