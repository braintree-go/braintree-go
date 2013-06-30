package braintree

import (
	"testing"
)

func TestCreateCreditCard(t *testing.T) {
	customer, err := testGateway.Customer().Create(&Customer{})
	if err != nil {
		t.Fatal(err)
	}
	card, err := testGateway.CreditCard().Create(&CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: true,
		},
	})

	if err != nil {
		t.Fatal(err)
	}
	if card.Token == "" {
		t.Fatal("invalid token")
	}
}

func TestCreateCreditCardInvalidInput(t *testing.T) {
	card, err := testGateway.CreditCard().Create(&CreditCard{
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
	})

	t.Log(card)

	// This test should fail because customer id is required
	if err == nil {
		t.Fail()
	}

	// TODO: validate fields
}

func TestFindCreditCard(t *testing.T) {
	customer, err := testGateway.Customer().Create(&Customer{})
	if err != nil {
		t.Fatal(err)
	}
	card, err := testGateway.CreditCard().Create(&CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: true,
		},
	})

	t.Log(card)

	if err != nil {
		t.Fatal(err)
	}
	if card.Token == "" {
		t.Fatal("invalid token")
	}

	card2, err := testGateway.CreditCard().Find(card.Token)

	t.Log(card2)

	if err != nil {
		t.Fatal(err)
	}
	if card2.Token != card.Token {
		t.Fatal("tokens do not match")
	}
}

func TestFindCreditCardBadData(t *testing.T) {
	card, err := testGateway.CreditCard().Find("invalid_token")

	t.Log(card)

	if err == nil {
		t.Fail()
	}
}

func TestSaveCreditCardWithVenmoSDKPaymentMethodCode(t *testing.T) {
	customer, err := testGateway.Customer().Create(&Customer{})
	if err != nil {
		t.Fatal(err)
	}
	card, err := testGateway.CreditCard().Create(&CreditCard{
		CustomerId:                customer.Id,
		VenmoSDKPaymentMethodCode: "stub-" + testCreditCards["visa"].Number,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !card.VenmoSDK {
		t.Fatal("venmo card not marked")
	}
}

func TestSaveCreditCardWithVenmoSDKSession(t *testing.T) {
	customer, err := testGateway.Customer().Create(&Customer{})
	if err != nil {
		t.Fatal(err)
	}
	card, err := testGateway.CreditCard().Create(&CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		Options: &CreditCardOptions{
			VenmoSDKSession: "stub-session",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !card.VenmoSDK {
		t.Fatal("venmo card not marked")
	}
}
