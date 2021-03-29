// +build unit

package braintree

import (
	"encoding/xml"
	"testing"
)

func TestApplePayCard_MarshalXML(t *testing.T) {
	card := ApplePayCard{
		ExpirationMonth: "10",
		ExpirationYear:  "22",
		ECI:             "07",
		Cryptogram:      "testCardCryptogram",
		Number:          "4111111111111111",
		CardholderName:  "Test User",
	}

	cardBytes, err := xml.Marshal(card)
	if err != nil {
		t.Fatalf("%v", err)
	}

	expected := `<apple-pay-card><expiration-month>10</expiration-month><expiration-year>22</expiration-year><eci-indicator>07</eci-indicator><cryptogram>testCardCryptogram</cryptogram><number>4111111111111111</number><cardholder-name>Test User</cardholder-name></apple-pay-card>`

	if string(cardBytes) != expected {
		t.Errorf("Marshalled xml got [%s], want [%s]", cardString, expected)
	}
}