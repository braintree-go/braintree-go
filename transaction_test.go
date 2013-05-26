package braintree

import (
	"testing"
)

func TestToXML(t *testing.T) {
	tx := Transaction{
		Type:   "sale",
		Amount: 100,
		CreditCard: &CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}
	expectedXML := "<transaction><type>sale</type><amount>100</amount><credit-card><number>" + TestCreditCards["visa"].Number + "</number><expiration-date>05/14</expiration-date></credit-card></transaction>"

	xmlBytes, err := tx.ToXML()
	generatedXML := string(xmlBytes)
	if err != nil {
		t.Errorf("Got error while generating XML: " + err.Error())
	} else if generatedXML != expectedXML {
		t.Errorf("Got incorrect XML: " + generatedXML)
	}
}

func TestSaleErrorHandling(t *testing.T) {
	gateway := TransactionGateway{blowUpGateway{}}
	response, err := gateway.Sale(Transaction{})

	if response.IsSuccess() {
		t.Errorf("Sale response was successful when the server 500'd")
	} else if err.Error() != "The server blew up!" {
		t.Errorf("Sale returned wrong error when the server 500'd")
	}
}

func TestSaleBadInputHandling(t *testing.T) {
	gateway = TransactionGateway{badInputGateway{}}
	response, err := gateway.Sale(Transaction{})

	if response.IsSuccess() {
		t.Errorf("Sale response was successful when the server 422'd")
	} else if err != nil {
		t.Errorf("Sale returned an error that should be handled on response object")
	} else if response.GetMessage() != "Card Issuer Declined CVV" {
		t.Errorf("Sale returned wrong error message")
	}
}

func TestFindErrorHandling(t *testing.T) {
	gateway := TransactionGateway{notFoundGateway{}}
	response, err := gateway.Find("bogusId")

	if response.IsSuccess() {
		t.Errorf("Find response was successful when the server 404'd")
	} else if err.Error() != "A transaction with that ID could not be found" {
		t.Errorf("Find returned wrong error when the server 404'd")
	}
}
