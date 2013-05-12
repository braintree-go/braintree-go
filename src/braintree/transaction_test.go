package braintree

import "testing"

func TestTransactionRequestXML(t *testing.T) {
	tx := Transaction{
		Amount: 100,
		CreditCard: CreditCard{
			Number:         TestCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}
	request := NewTransactionRequest(tx)
	expectedXML := "<transaction><amount>100</amount><credit-card><number>" + TestCards["visa"].Number + "</number><expiration-date>05/14</expiration-date></credit-card></transaction>"

	xmlBytes, err := request.ToXML()
	generatedXML := string(xmlBytes)
	if err != nil {
		t.Errorf("Got error while generating XML: " + err.Error())
	} else if generatedXML != expectedXML {
		t.Errorf("Got incorrect XML: " + generatedXML)
	}
}
