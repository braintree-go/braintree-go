package braintree

import (
	"testing"
)

func TestTransactionRequestXML(t *testing.T) {
	tx := Transaction{
		Type:   "sale",
		Amount: 100,
		CreditCard: CreditCard{
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
